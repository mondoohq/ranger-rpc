// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package swagger

import (
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/go-openapi/spec"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func New() *swaggerGen {
	return &swaggerGen{&pgs.ModuleBase{}}
}

type swaggerGen struct {
	*pgs.ModuleBase
}

func (pc *swaggerGen) Name() string {
	return "Ranger RPC Swagger 2.0"
}

func (pc *swaggerGen) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		ctx := pc.Push(f.Name().String())
		goctx := pgsgo.InitContext(ctx.Parameters())

		pkgname := goctx.PackageName(f).String()
		pc.Debugf("processing file %s with pkg %s", f.Name(), pkgname)

		// do not create a file if no service is part of the definition
		if len(f.Services()) == 0 {
			continue
		}

		// TODO: derive version from go package version
		sw := NewSwaggerEndpoint(f.Name().String(), "1.0")

		for _, service := range f.Services() {

			// collect import paths for each protobuf method
			for _, m := range service.Methods() {

				protoPkg := f.Package().ProtoName().Split()

				// remove version pkg if its the last
				if len(protoPkg) > 1 {
					versionPackagePart := regexp.MustCompile(`^v(\d*)$`)
					last := protoPkg[len(protoPkg)-1]
					if versionPackagePart.MatchString(last) {
						protoPkg = protoPkg[:len(protoPkg)-1]
					}
				}

				action := fmt.Sprintf("%s.%s.%s", strings.Join(protoPkg, "."), service.Name(), m.Name())
				pc.Debug(action)

				methodPath := fmt.Sprintf("/%s/%s", service.Name().String(), m.Name().String())
				pc.Debug(methodPath)

				inputDefinitionName := fmt.Sprintf("%s.%s.%s", strings.Join(protoPkg, "."), service.Name(), m.Input().Name())
				inputRef := fmt.Sprintf("#/definitions/%s", inputDefinitionName)
				sw.AddProtoMessage(inputDefinitionName, m.Input())
				pc.Debug(inputRef)

				outputDefinitionName := fmt.Sprintf("%s.%s.%s", strings.Join(protoPkg, "."), service.Name(), m.Output().Name())
				outputRef := fmt.Sprintf("#/definitions/%s", outputDefinitionName)
				sw.AddProtoMessage(outputDefinitionName, m.Output())
				pc.Debug(outputRef)

				sw.AddProtoMethod(methodPath, m.Name().String(), inputRef, outputRef, m)
			}

		}

		// generate json output
		fp := pgs.FilePath(ctx.OutputPath())
		b, _ := json.MarshalIndent(sw, "", "  ")
		pc.AddGeneratorFile(fp.SetBase(f.Name().String()).SetExt(".swagger.json").String(), string(b))
		pc.Pop()
	}
	return pc.Artifacts()
}

func NewSwaggerEndpoint(filename string, version string) *SwaggerApi {
	sw := &SwaggerApi{}
	sw.Swagger.Swagger = "2.0"
	sw.Schemes = []string{"http", "https"}
	sw.Produces = []string{"application/json"}
	sw.Consumes = sw.Produces
	sw.Definitions = spec.Definitions{}
	sw.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:   path.Base(filename),
			Version: version,
		},
	}
	sw.Paths = &spec.Paths{
		Paths: make(map[string]spec.PathItem),
	}
	return sw
}

type SwaggerApi struct {
	spec.Swagger
}

func (sw *SwaggerApi) AddProtoMessage(definitionName string, message pgs.Message) {

	schemaProps := make(map[string]spec.Schema)

	fields := message.Fields()
	for i := range fields {
		f := fields[i]

		fieldName := f.Name().String()
		fieldTitle := f.Name().String()
		fieldDescription := strings.TrimSpace(f.SourceCodeInfo().LeadingComments())
		fieldType, fieldFormat := swaggerFieldType(f)

		schemaProps[fieldName] = spec.Schema{
			SchemaProps: spec.SchemaProps{
				Title:       fieldTitle,
				Description: fieldDescription,
				Type:        spec.StringOrArray([]string{fieldType}),
				Format:      fieldFormat,
				//Items: &spec.SchemaOrArray{
				//	Schema: &fieldSchema,
				//},
			},
		}
	}

	comments := message.SourceCodeInfo().LeadingComments()

	sw.Swagger.Definitions[definitionName] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Title:       title(comments),
			Description: strings.TrimSpace(comments),
			Type:        spec.StringOrArray([]string{"object"}),
			Properties:  schemaProps,
		},
	}
}

func title(c string) string {
	lines := strings.Split(c, "\n")
	if len(lines) > 1 {
		return strings.TrimSpace(lines[0])
	}
	return ""
}

// swaggerFieldType extracts the open api v2 types from the proto message
// see https://swagger.io/specification/v2/#data-types
// see https://developers.google.com/protocol-buffers/docs/proto3#scalar
// - type: the swagger data type
// - format: optional more detailed information about the type
func swaggerFieldType(f pgs.Field) (string, string) {
	var fieldType, fieldFormat string

	switch f.Type().ProtoType() {
	case pgs.Int32T, pgs.SFixed32, pgs.SInt32:
		fieldType = "integer"
		fieldFormat = "int32"
	case pgs.Int64T, pgs.SFixed64, pgs.SInt64:
		fieldType = "integer"
		fieldFormat = "int64"
	case pgs.Fixed32T, pgs.UInt32T:
		fieldType = "integer"
		fieldFormat = "uint32"
	case pgs.Fixed64T, pgs.UInt64T:
		fieldType = "integer"
		fieldFormat = "uint64"
	case pgs.DoubleT:
		fieldType = "number"
		fieldFormat = "double"
	case pgs.FloatT:
		fieldType = "number"
		fieldFormat = "float"
	case pgs.BoolT:
		fieldType = "boolean"
		fieldFormat = ""
	case pgs.StringT:
		fieldType = "string"
		fieldFormat = ""
	case pgs.BytesT:
		fieldType = "string"
		fieldFormat = "byte"
	case pgs.EnumT:
		// https://swagger.io/docs/specification/2-0/enums/
		// NOTE: enums are not natively supported, we do not add the yaml definition yet
		fieldType = "string"
	case pgs.MessageT:
		fieldType = "object"
		// TODO: we need to generated the nested type here
	case pgs.GroupT:
		// NOTE: GroupT is also not properly supported by pgs
		panic(f.Name().String() + ": proto field type GroupT not supported")

	}

	return fieldType, fieldFormat
}

func (sw *SwaggerApi) AddProtoMethod(methodPath string, id string, inputRef string, outputRef string, m pgs.Method) {

	comments := m.SourceCodeInfo().LeadingComments()

	sw.Paths.Paths[methodPath] = spec.PathItem{
		PathItemProps: spec.PathItemProps{
			Post: &spec.Operation{
				OperationProps: spec.OperationProps{
					ID:          id,
					Summary:     title(comments),
					Description: strings.TrimSpace(comments),
					Parameters: []spec.Parameter{
						spec.Parameter{
							ParamProps: spec.ParamProps{
								Name:     "body",
								In:       "body",
								Required: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: spec.MustCreateRef(inputRef),
									},
								},
							},
						},
					},
					Responses: &spec.Responses{
						ResponsesProps: spec.ResponsesProps{
							StatusCodeResponses: map[int]spec.Response{
								200: spec.Response{
									ResponseProps: spec.ResponseProps{
										Schema: &spec.Schema{
											SchemaProps: spec.SchemaProps{
												Ref: spec.MustCreateRef(outputRef),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
