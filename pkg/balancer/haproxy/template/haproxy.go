/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package template

// HAProxy configuration template
const HAProxy = `
# Generated HAProxy
{{ with .Global}}
  global
  log 127.0.0.1   local0
  log 127.0.0.1   local1 notice
  maxconn 20000
  tune.ssl.default-dh-param 2048

  ssl-default-bind-options no-sslv3 no-tls-tickets
  ssl-default-bind-ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA

  ssl-default-server-options no-sslv3 no-tls-tickets
  ssl-default-server-ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA
  daemon
{{ end }}

{{ with .Defaults }}
defaults
  mode    http
  option forwardfor
  option http-keep-alive
  timeout connect 5000
  maxconn 20000
  timeout client  50000
  timeout server  50000
  timeout http-keep-alive 100s
  timeout tunnel        3600s
  monitor-uri /health_check
{{ end }}


frontend stats
  bind *:8076
  mode http
  stats enable
  stats hide-version
  stats uri /
  stats refresh 5s
  timeout client 60s
  timeout server 60s
  timeout connect 60s

{{$certs_dir:= .CertsDir }}

{{ range .Frontends }}
frontend {{ .Name }}
  {{ with .Bind }}
  bind {{ .IP }}:{{ .Port }}{{ if .IsTLS }} ssl {{ range .Certs }} crt {{$certs_dir}}/{{.Name}}.pem {{ end }} {{ end }}{{ end }}
  {{ if .DefaultBackend.Backend }}
    default_backend {{ .DefaultBackend.Backend }}
  {{end}}
  acl geo_country_code_us_unknown hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-UNKNOWN.txt
  acl geo_country_code_us_ak hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-AK.txt
  acl geo_country_code_us_al hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-AL.txt
  acl geo_country_code_us_ar hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-AR.txt
  acl geo_country_code_us_az hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-AZ.txt
  acl geo_country_code_us_ca hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-CA.txt
  acl geo_country_code_us_co hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-CO.txt
  acl geo_country_code_us_ct hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-CT.txt
  acl geo_country_code_us_dc hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-DC.txt
  acl geo_country_code_us_de hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-DE.txt
  acl geo_country_code_us_fl hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-FL.txt
  acl geo_country_code_us_ga hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-GA.txt
  acl geo_country_code_us_hi hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-HI.txt
  acl geo_country_code_us_ia hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-IA.txt
  acl geo_country_code_us_id hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-ID.txt
  acl geo_country_code_us_il hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-IL.txt
  acl geo_country_code_us_in hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-IN.txt
  acl geo_country_code_us_ks hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-KS.txt
  acl geo_country_code_us_ky hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-KY.txt
  acl geo_country_code_us_la hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-LA.txt
  acl geo_country_code_us_ma hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MA.txt
  acl geo_country_code_us_md hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MD.txt
  acl geo_country_code_us_me hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-ME.txt
  acl geo_country_code_us_mi hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MI.txt
  acl geo_country_code_us_mn hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MN.txt
  acl geo_country_code_us_mo hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MO.txt
  acl geo_country_code_us_ms hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MS.txt
  acl geo_country_code_us_mt hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-MT.txt
  acl geo_country_code_us_nc hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NC.txt
  acl geo_country_code_us_nd hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-ND.txt
  acl geo_country_code_us_ne hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NE.txt
  acl geo_country_code_us_nh hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NH.txt
  acl geo_country_code_us_nj hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NJ.txt
  acl geo_country_code_us_nm hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NM.txt
  acl geo_country_code_us_nv hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NV.txt
  acl geo_country_code_us_ny hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-NY.txt
  acl geo_country_code_us_oh hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-OH.txt
  acl geo_country_code_us_ok hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-OK.txt
  acl geo_country_code_us_or hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-OR.txt
  acl geo_country_code_us_pa hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-PA.txt
  acl geo_country_code_us_ri hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-RI.txt
  acl geo_country_code_us_sc hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-SC.txt
  acl geo_country_code_us_sd hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-SD.txt
  acl geo_country_code_us_tn hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-TN.txt
  acl geo_country_code_us_tx hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-TX.txt
  acl geo_country_code_us_ut hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-UT.txt
  acl geo_country_code_us_va hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-VA.txt
  acl geo_country_code_us_vt hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-VT.txt
  acl geo_country_code_us_wa hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-WA.txt
  acl geo_country_code_us_wi hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-WI.txt
  acl geo_country_code_us_wv hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-WV.txt
  acl geo_country_code_us_wy hdr_ip(X-Forwarded-For,1) -f /opt/stackpoint/balancer/geodb/US-WY.txt

  http-request set-header X-GEOIP_REGION XX if geo_country_code_us_unknown
  http-request set-header X-GEOIP_REGION AK if geo_country_code_us_ak
  http-request set-header X-GEOIP_REGION AL if geo_country_code_us_al
  http-request set-header X-GEOIP_REGION AR if geo_country_code_us_ar
  http-request set-header X-GEOIP_REGION AZ if geo_country_code_us_az
  http-request set-header X-GEOIP_REGION CA if geo_country_code_us_ca
  http-request set-header X-GEOIP_REGION CO if geo_country_code_us_co
  http-request set-header X-GEOIP_REGION CT if geo_country_code_us_ct
  http-request set-header X-GEOIP_REGION DC if geo_country_code_us_dc
  http-request set-header X-GEOIP_REGION DE if geo_country_code_us_de
  http-request set-header X-GEOIP_REGION FL if geo_country_code_us_fl
  http-request set-header X-GEOIP_REGION GA if geo_country_code_us_ga
  http-request set-header X-GEOIP_REGION HI if geo_country_code_us_hi
  http-request set-header X-GEOIP_REGION IA if geo_country_code_us_ia
  http-request set-header X-GEOIP_REGION ID if geo_country_code_us_id
  http-request set-header X-GEOIP_REGION IL if geo_country_code_us_il
  http-request set-header X-GEOIP_REGION IN if geo_country_code_us_in
  http-request set-header X-GEOIP_REGION KS if geo_country_code_us_ks
  http-request set-header X-GEOIP_REGION KY if geo_country_code_us_ky
  http-request set-header X-GEOIP_REGION LA if geo_country_code_us_la
  http-request set-header X-GEOIP_REGION MA if geo_country_code_us_ma
  http-request set-header X-GEOIP_REGION MD if geo_country_code_us_md
  http-request set-header X-GEOIP_REGION ME if geo_country_code_us_me
  http-request set-header X-GEOIP_REGION MI if geo_country_code_us_mi
  http-request set-header X-GEOIP_REGION MN if geo_country_code_us_mn
  http-request set-header X-GEOIP_REGION MO if geo_country_code_us_mo
  http-request set-header X-GEOIP_REGION MS if geo_country_code_us_ms
  http-request set-header X-GEOIP_REGION MT if geo_country_code_us_mt
  http-request set-header X-GEOIP_REGION NC if geo_country_code_us_nc
  http-request set-header X-GEOIP_REGION ND if geo_country_code_us_nd
  http-request set-header X-GEOIP_REGION NE if geo_country_code_us_ne
  http-request set-header X-GEOIP_REGION NH if geo_country_code_us_nh
  http-request set-header X-GEOIP_REGION NJ if geo_country_code_us_nj
  http-request set-header X-GEOIP_REGION NM if geo_country_code_us_nm
  http-request set-header X-GEOIP_REGION NV if geo_country_code_us_nv
  http-request set-header X-GEOIP_REGION NY if geo_country_code_us_ny
  http-request set-header X-GEOIP_REGION OH if geo_country_code_us_oh
  http-request set-header X-GEOIP_REGION OK if geo_country_code_us_ok
  http-request set-header X-GEOIP_REGION OR if geo_country_code_us_or
  http-request set-header X-GEOIP_REGION PA if geo_country_code_us_pa
  http-request set-header X-GEOIP_REGION RI if geo_country_code_us_ri
  http-request set-header X-GEOIP_REGION SC if geo_country_code_us_sc
  http-request set-header X-GEOIP_REGION SD if geo_country_code_us_sd
  http-request set-header X-GEOIP_REGION TN if geo_country_code_us_tn
  http-request set-header X-GEOIP_REGION TX if geo_country_code_us_tx
  http-request set-header X-GEOIP_REGION UT if geo_country_code_us_ut
  http-request set-header X-GEOIP_REGION VA if geo_country_code_us_va
  http-request set-header X-GEOIP_REGION VT if geo_country_code_us_vt
  http-request set-header X-GEOIP_REGION WA if geo_country_code_us_wa
  http-request set-header X-GEOIP_REGION WI if geo_country_code_us_wi
  http-request set-header X-GEOIP_REGION WV if geo_country_code_us_wv
  http-request set-header X-GEOIP_REGION WY if geo_country_code_us_wy

  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_unknown
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ak
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_al
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ar
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_az
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ca
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_co
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ct
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_dc
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_de
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_fl
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ga
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_hi
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ia
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_id
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_il
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_in
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ks
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ky
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_la
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ma
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_md
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_me
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_mi
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_mn
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_mo
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ms
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_mt
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_nc
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_nd
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ne
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_nh
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_nj
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_nm
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_nv
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ny
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_oh
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ok
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_or
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_pa
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ri
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_sc
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_sd
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_tn
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_tx
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_ut
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_va
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_vt
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_wa
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_wi
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_wv
  http-request set-header X-GEOIP_COUNTRY_CODE US if geo_country_code_us_wy
  acl HAS_CF_IPCOUNTRY req.fhdr(CF-IPCountry) -m found
  http-request set-header X-GEOIP_COUNTRY_CODE %[req.hdr(CF-IPCountry)] if HAS_CF_IPCOUNTRY

  {{ range .ACLs }}
    acl {{ .Name }} {{.Content}}
  {{end}}
  {{ range .UseBackendsByPrio }}
    use_backend {{ .Backend }} if {{ range .ACLs }}{{ .Name }} {{end}}
  {{end}}
{{ end }}

{{range $name, $be := .Backends}}
backend {{$name}}{{ range $sname, $server := .Servers}}
  mode http
  balance leastconn
  option http-keep-alive
  option forwardfor
  option httpchk HEAD /status HTTP/1.0
  cookie SRV_ID prefix
  server {{ $sname }} {{ $server.Address }}:{{ $server.Port }} check inter {{ $server.CheckInter}} fall 3 rise 2 {{end}}
{{end}}
`
