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
	//go:embed templates/gofile.template
	embeddedTemplateGofile string

	//go:embed templates/service_client.template
	embeddedTemplateServiceClient string

	//go:embed templates/service_server.template
	embeddedTemplateServiceServer string

	templateGoRangerFile    *template.Template
	templateGoServiceClient *template.Template
	templateGoServiceServer *template.Template
)

func noescape(str string) template.HTML {
	return template.HTML(str)
}

func init() {
	var fn = template.FuncMap{
		"noescape": noescape,
		"gotype":   goMessageType,
	}

	// load file content
	templateGoRangerFile = template.New("file")
	templateGoRangerFile.Funcs(fn)
	tmpl, err := templateGoRangerFile.Parse(embeddedTemplateGofile)
	if err != nil {
		panic(err)
	}
	templateGoRangerFile = tmpl

	// load service client content
	templateGoServiceClient = template.New("service_client")
	templateGoServiceClient.Funcs(fn)
	tmpl, err = templateGoServiceClient.Parse(embeddedTemplateServiceClient)
	if err != nil {
		panic(err)
	}
	templateGoServiceClient = tmpl

	// load service server content
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

			// render service server
			serviceServer, err := fc.renderServiceServer(goServiceRenderOpts{
				Pkg:     pkg,
				Service: service,
			})
			fc.CheckErr(err, "unable to render ", service, " server to proto")
			servicesRendered += serviceServer
		}

		// generate complete file, services are included
		fileContent, err := fc.renderFile(goFileRenderOpts{
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

func (fc *rangerc) renderFile(renderOpts goFileRenderOpts) (string, error) {
	return fc.render(templateGoRangerFile, renderOpts)
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
