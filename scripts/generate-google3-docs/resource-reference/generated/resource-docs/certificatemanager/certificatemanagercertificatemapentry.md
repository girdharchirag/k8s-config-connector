{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}CertificateManagerCertificateMapEntry{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Certificate Manager</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/certificate-manager/docs/">/certificate-manager/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1.projects.locations.certificateMaps.certificateMapEntries</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/certificate-manager/docs/reference/certificate-manager/rest/v1/projects.locations.certificateMaps.certificateMapEntries">/certificate-manager/docs/reference/certificate-manager/rest/v1/projects.locations.certificateMaps.certificateMapEntries</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpcertificatemanagercertificatemapentry<br>gcpcertificatemanagercertificatemapentries<br>certificatemanagercertificatemapentry</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>certificatemanager.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>certificatemanagercertificatemapentries.certificatemanager.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


<tr>
<td>{{product_name_short}} Default Average Reconcile Interval In Seconds</td>
<td>600</td>
</tr>
</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/state-into-spec</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
certificatesRefs:
- external: string
  name: string
  namespace: string
description: string
hostname: string
mapRef:
  external: string
  name: string
  namespace: string
matcher: string
projectRef:
  external: string
  name: string
  namespace: string
resourceID: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>certificatesRefs</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>certificatesRefs[]</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}A set of Certificates defines for the given hostname.
There can be defined up to fifteen certificates in each Certificate Map Entry.
Each certificate must match pattern projects/*/locations/*/certificates/*.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>certificatesRefs[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: string of the format `projects/{{project}}/locations/{{location}}/certificates/{{value}}`, where {{value}} is the `name` field of a `CertificateManagerCertificate` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>certificatesRefs[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>certificatesRefs[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A human-readable description of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>hostname</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
selecting a proper certificate.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>mapRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}A map entry that is inputted into the certificate map.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>mapRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of a `CertificateManagerCertificateMap` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>mapRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>mapRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>matcher</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. A predefined matcher for particular cases, other than SNI selection.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The project that this resource belongs to.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of a `Project` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>projectRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
createTime: string
observedGeneration: integer
state: string
updateTime: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>createTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A serving state of this Certificate Map Entry.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>updateTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
```yaml
# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: certificatemanager.cnrm.cloud.google.com/v1beta1
kind: CertificateManagerCertificateMapEntry
metadata:
  name: certificatemanagercertificatemapentry-sample
spec:
  description: sample certificate map entry
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID.
    external: ${PROJECT_ID?}
  matcher: PRIMARY
  certificatesRefs:
    - name: certificatemanagercertificatemapentry-dep
  mapRef:
    name: certificatemanagercertificatemapentry-dep
---
apiVersion: certificatemanager.cnrm.cloud.google.com/v1beta1
kind: CertificateManagerCertificate
metadata:
  name: certificatemanagercertificatemapentry-dep
spec:
  location : global
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID.
    external: ${PROJECT_ID?}
  selfManaged:
    pemCertificate: |-
      -----BEGIN CERTIFICATE-----
      MIIDDzCCAfegAwIBAgIUDOiCLH9QNMMYnjPZVf4VwO9blsEwDQYJKoZIhvcNAQEL
      BQAwFjEUMBIGA1UEAwwLZXhhbXBsZS5jb20wIBcNMjIwODI0MDg0MDUxWhgPMzAy
      MTEyMjUwODQwNTFaMBYxFDASBgNVBAMMC2V4YW1wbGUuY29tMIIBIjANBgkqhkiG
      9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvOT925GG4lKV9HvAHsbecMhGPAqjhVRC26iZ
      UJC8oSWOu95lWJSX5ZhbiF6Nz192wDGV/VAh3Lxj8RYtcn75eDxQKTcKouDld+To
      CGIStPFWbR6rbysLuZqFVEXVOTvp2QIegInfrvnGC4j7Qpic7zrFB9HzJx+0HpeE
      yO4gkdzJfEK/gMmolUgJrKX59o+0+Rj+Jq3EtcQxL1fVBVJSx0NvpoR1eYpnHMr/
      rJKZkUUZ2xE86hrtpiP6OEYQTi00rmf4GnZF5QfGGD0xuoQXtR7Tu+XhKibXIhxc
      D4RzPLX1QS040PXvmMPLDb4YlUQ6V3Rs42JDvkkDwIMXZvn8awIDAQABo1MwUTAd
      BgNVHQ4EFgQURuo1CCZZAUv7xi02f2nC5tRbf18wHwYDVR0jBBgwFoAURuo1CCZZ
      AUv7xi02f2nC5tRbf18wDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOC
      AQEAqx3tDxurnYr9EUPhF5/LlDPYM+VI7EgrKdRnuIqUlZI0tm3vOGME0te6dBTC
      YLNaHLW3m/4Tm4M2eg0Kpz6CxJfn3109G31dCi0xwzSDHf5TPUWvqIVhq5WRgMIf
      n8KYBlQSmqdJBRztUIQH/UPFnSbxymlS4s5qwDgTH5ag9EEBcnWsQ2LZjKi0eqve
      MaqAvvB+j8RGZzYY4re94bSJI42zIZ6nMWPtXwRuDc30xl/u+E0jWIgWbPwSd6Km
      3wnJnGiU2ezPGq3zEU+Rc39VVIFKQpciNeYuF3neHPJvYOf58qW2Z8s0VH0MR1x3
      3DoO/e30FIr9j+PRD+s5BPKF2A==
      -----END CERTIFICATE-----
    pemPrivateKey:
      valueFrom:
        secretKeyRef:
          name: certificatemanagercertificatemapentry-dep
          key: privateKey
---
apiVersion: certificatemanager.cnrm.cloud.google.com/v1beta1
kind: CertificateManagerCertificateMap
metadata:
  name: certificatemanagercertificatemapentry-dep
spec:
  projectRef:
    # Replace ${PROJECT_ID?} with your project ID.
    external: ${PROJECT_ID?}
---
apiVersion: v1
kind: Secret
metadata:
  name: certificatemanagercertificatemapentry-dep
stringData:
  privateKey: |
    -----BEGIN PRIVATE KEY-----
    MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC85P3bkYbiUpX0
    e8Aext5wyEY8CqOFVELbqJlQkLyhJY673mVYlJflmFuIXo3PX3bAMZX9UCHcvGPx
    Fi1yfvl4PFApNwqi4OV35OgIYhK08VZtHqtvKwu5moVURdU5O+nZAh6Aid+u+cYL
    iPtCmJzvOsUH0fMnH7Qel4TI7iCR3Ml8Qr+AyaiVSAmspfn2j7T5GP4mrcS1xDEv
    V9UFUlLHQ2+mhHV5imccyv+skpmRRRnbETzqGu2mI/o4RhBOLTSuZ/gadkXlB8YY
    PTG6hBe1HtO75eEqJtciHFwPhHM8tfVBLTjQ9e+Yw8sNvhiVRDpXdGzjYkO+SQPA
    gxdm+fxrAgMBAAECggEAV4/A24TQpV4KFBw/WSTvnRFBeXinB1mhamhztWR6hCrA
    SPcVPKQY632eRI8sJmpGxl3V/Ogl4khT/cA9jfstEl7G++v/WrRsupCaPLSVnlnX
    KdsTNgOauk1WK9P5PMA4rPcuA4Cl91riQpubeWn8KWsxRWg90i+Ak8PB8lBsOaB1
    QzjigWlrRWSpodaw0MBIMZFDL2BYK8HEr+wyATYIyGvDQc9zCnMQIQIZyEPYepLO
    04Dw17YcjgnoJ5gLAFiTvDrCpTMewud1RQzvW5TAvG2piw34sf3QMGPM7aXNrfuZ
    4ZPC/MwVQgq9Nc+jeDsjApQmJKJ+3a8OdIPU89ArTQKBgQDCpHHQe1RzpHmIx47/
    9N5r+NPBhh8flDYmvgi6zPeBfrAaLWhidS8c7Voa6HwvMxbhryDEvc0YqI3vllfy
    xnRF+DfSryozW0gjrkXDGoOzqOJ3EuQwLSJnyX6La2lmufqsRFazwYJ5sxcjoGHK
    /sbwZkIUj1ejuH44ve+ZJQFfpwKBgQD4cLJrJhqImUDhHZRx9jBvxyeHy/RjmHK6
    70xQVDi9ZqeExHwtoSbolhXKLB1RtBnw+t5Csy7IDNBDsbUg9fXU8KyCTIdmsyws
    bDb5hdKsUF76rkKzlpttiXMRVWGS3CMKWahBpnL3lFB3tdtmskemkBTXVn4VgKAH
    xk9XnZ11nQKBgDbQSJ0FnkrSzscOK984/ko50Kh3NNyXyIgwjBTPFASLwNweXX8c
    sR/cV7usLQy9vnvf7cJ6EQAYt5/5Httnt+bceBwE6EV+N1qVAWBoXx6BOQV/dHN8
    wmun+tMYdJ5RUZ6hwCjvHedX3/RQfjnEdhHNOl6/31Zj5mfkVU0zdqeRAoGAcvIh
    erXMfPr7K6y16+xOCMmKHqhc0F/OZXMmSdxNzEPcqe8GzU3MZLxcJIg4oH7FqdtI
    Tm/86w4Spd9owHFMZlNcXYTu+LNZcsw2u0gRayxcZXuO3OyHySxZEuIAHSTBCZ7l
    3EoY0zfJ6zk249MEl6n+GouoFmbGpBI6z3zbR3kCgYEAlCNZVH4uJrP5beTOZTTR
    VJRk7BXvEC6HsM140YtIN7NHy2GtzrgmmY/ZAFB/hX8Ft4ex2MxbIp3hvxroTqGn
    bfu7uv97NoPQqbjtc3Mz8h2IaXTVDUnWYY5gDu6rM2w+Z75/sWIGiTWrsdYX4ohb
    ujngzJ7Ew7GgKSboj6mtlVM=
    -----END PRIVATE KEY-----
```


Note: If you have any trouble with instantiating the resource, refer to <a href="/config-connector/docs/troubleshooting">Troubleshoot Config Connector</a>.

{% endblock %}