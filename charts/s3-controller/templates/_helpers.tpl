{{- define "s3-controller.name" -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}