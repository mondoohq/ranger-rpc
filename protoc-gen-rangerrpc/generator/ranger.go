package generator

import (
	"bufio"
	"bytes"
	_ "embed"
	"html/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

var (
	//go:embed templates/go_client_server_file.template
	embeddedTemplateGofileClientServer string

	//go:embed templates/go_client_file.template
	embeddedTemplateGofileClient string

	//go:embed templates/service_client.template
	embeddedTemplateServiceClient string

	//go:embed templates/service_server.template
	embeddedTemplateServiceServer string

	templateGoClientAndServerRangerFile *template.Template
	templateGoClientRangerFile          *template.Template
	templateGoServiceClient             *template.Template
	templateGoServiceServer             *template.Template
)

func noescape(str string) template.HTML {
	return template.HTML(str)
}

func init() {
	var fn = template.FuncMap{
		"noescape": noescape,
		"gotype":   goMessageType,
	}

	// load go client and server file template content
	templateGoClientAndServerRangerFile = template.New("file_client_server")
	templateGoClientAndServerRangerFile.Funcs(fn)
	tmpl, err := templateGoClientAndServerRangerFile.Parse(embeddedTemplateGofileClientServer)
	if err != nil {
		panic(err)
	}
	templateGoClientAndServerRangerFile = tmpl

	// load go client and server file template content
	templateGoClientRangerFile = template.New("file_client")
	templateGoClientRangerFile.Funcs(fn)
	tmpl, err = templateGoClientRangerFile.Parse(embeddedTemplateGofileClient)
	if err != nil {
		panic(err)
	}
	templateGoClientRangerFile = tmpl

	// load service client template content
	templateGoServiceClient = template.New("service_client")
	templateGoServiceClient.Funcs(fn)
	tmpl, err = templateGoServiceClient.Parse(embeddedTemplateServiceClient)
	if err != nil {
		panic(err)
	}
	templateGoServiceClient = tmpl

	// load service server template content
	templateGoServiceServer = template.New("service_server")
	templateGoServiceServer.Funcs(fn)
	tmpl, err = templateGoServiceServer.Parse(embeddedTemplateServiceServer)
	if err != nil {
		panic(err)
	}
	templateGoServiceServer = tmpl
}

func New() *rangerc {
	return &rangerc{&pgs.ModuleBase{}}
}

type rangerc struct {
	*pgs.ModuleBase
}

func (f *rangerc) Name() string {
	return "ranger"
}

func (fc *rangerc) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	clientOnly := false
	v, ok := fc.Parameters()["client-only"]
	if ok && v == "true" {
		clientOnly = true
	}

	for _, f := range targets {
		ctx := fc.Push(f.Name().String())
		goctx := pgsgo.InitContext(ctx.Parameters())

		pkgname := goctx.PackageName(f).String()
		fc.Debugf("processing file %s with pkg %s", f.Name(), pkgname)

		// do not create a file if no service is part of the definition
		if len(f.Services()) == 0 {
			continue
		}

		pkg := f.Package()
		servicesRendered := ""
		imports := make(map[string]bool)

		// iterate over each service
		for _, service := range f.Services() {

			// collect import paths for each protobuf method
			for _, m := range service.Methods() {
				fc.Debugf("signature %s, %s, %s", m.Name(), m.Input().Name(), m.Output().Name())

				if pkg != m.Input().Package() {
					inputPkg := goctx.ImportPath(m.Input()).String()
					fc.Debugf("found additional import in input: %s", inputPkg)
					imports[inputPkg] = true
				}

				if pkg != m.Output().Package() {
					outputPkg := goctx.ImportPath(m.Output()).String()
					fc.Debugf("found additional import in output: %s", outputPkg)
					imports[outputPkg] = true
				}
			}

			// render service client
			serviceClient, err := fc.renderServiceClient(goServiceRenderOpts{
				Pkg:     pkg,
				Service: service,
			})
			fc.CheckErr(err, "unable to render ", service, " client to proto")
			servicesRendered += serviceClient

			if !clientOnly {
				// render service server
				serviceServer, err := fc.renderServiceServer(goServiceRenderOpts{
					Pkg:     pkg,
					Service: service,
				})
				fc.CheckErr(err, "unable to render ", service, " server to proto")
				servicesRendered += serviceServer
			}
		}

		// generate complete file, services are included
		fileTemplate := templateGoClientAndServerRangerFile
		if clientOnly {
			fileTemplate = templateGoClientRangerFile
		}

		fileContent, err := fc.render(fileTemplate, goFileRenderOpts{
			Version: "version",
			Source:  f.Name().String(),
			Package: pkgname,
			Service: servicesRendered,
			Imports: imports,
		})
		fc.CheckErr(err, "unable to convert ", f.Name().String(), " to proto")

		fp := pgs.FilePath(ctx.OutputPath())
		fc.AddGeneratorFile(fp.SetBase(f.Name().String()).SetExt(".ranger.go").String(), fileContent)
		fc.Pop()
	}
	fc.Debugf("processing ranger generator completed")
	return fc.Artifacts()
}

// those are the input options for the gofile.templace
type goFileRenderOpts struct {
	Version string
	Source  string
	Package string
	Service string
	Imports map[string]bool
}

// those are the input options for the service.templace
type goServiceRenderOpts struct {
	Pkg     pgs.Package
	Service pgs.Service
}

func (fc *rangerc) renderServiceClient(renderOpts goServiceRenderOpts) (string, error) {
	return fc.render(templateGoServiceClient, renderOpts)
}

func (fc *rangerc) renderServiceServer(renderOpts goServiceRenderOpts) (string, error) {
	return fc.render(templateGoServiceServer, renderOpts)
}

// render a given go template
func (fc *rangerc) render(tmplate *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	err := tmplate.Execute(writer, data)
	if err != nil {
		return "", err
	}
	writer.Flush()
	return buf.String(), nil
}

// checks if the pkg is identical of the message type pkg. if the packages
// are not identical, it adds the
func goMessageType(pkg pgs.Package, m pgs.Message) string {
	goctx := pgsgo.InitContext(pgs.Parameters{})

	typeName := pgsgo.PGGUpperCamelCase(m.Name()).String()

	// if the method
	if pkg == m.Package() {
		return typeName
	}

	return goctx.PackageName(m.Package()).String() + "." + typeName
}
