package mock

/*
 * MIT License
 *
 * Copyright (c) 2024 Bamboo
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

import (
	"log"
	"sync"

	"github.com/GoSimplicity/AI-CloudOps/internal/model"
	"gorm.io/gorm"
)

type K8sClientMock struct {
	sync.RWMutex
	db *gorm.DB
}

func NewK8sClientMock(db *gorm.DB) *K8sClientMock {
	mock := &K8sClientMock{
		db: db,
	}

	mock.populateMockData()

	return mock
}

// populateMockData 插入多个模拟集群和相关数据到数据库
func (m *K8sClientMock) populateMockData() {
	m.Lock()
	defer m.Unlock()

	clusters := []model.K8sCluster{
		{
			Name:          "Cluster-1",
			NameZh:        "集群-1",
			UserID:        1,
			CpuRequest:    "100m",
			CpuLimit:      "200m",
			MemoryRequest: "256Mi",
			MemoryLimit:   "512Mi",
			Env:           "prod",
			Version:       "v1.32.0",
			ApiServerAddr: "https://api.cluster1.example.com",
			KubeConfigContent: `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCVENDQWUyZ0F3SUJBZ0lJSVVvMDZCMEplSFl3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TkRBNU16QXdPREEyTWpkYUZ3MHpOREE1TWpnd09ERXhNamRhTUJVeApFekFSQmdOVkJBTVRDbXQxWW1WeWJtVjBaWE13Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLCkFvSUJBUUREOW02ZWZQQytMazZ2clkwVEUyMVJieU9WZGMvVmR3c0FMZ011ci9zMlR0MzVYV2I4aFczT3lrQTEKNEVPYWQ5VmNrNExnK1VnMXdDV1ozMTNPbnFmRUFEQ25OWjFiRHJGbCs0Smhya2c4M0pUanlPZStqZmdNOXFWTQptczVoSjhnT3didmQ2WmdaOFl2bHowbHZGU3hEZGdHNXhIZGdqemZYU0FMRlhIS1hweVpWbHpqaWNWT1FNRWlrClM3clBoQUEyTnNLdDljeFVHUkI1OUMzN3poNks5MjdFclNJbUlKalZ0Ynl4MXJaVnBFSW1STGRPcFI2NWZlYnoKU2h6UlVwejVxTGMxZjl5TWlnWVRFY2tRZHFlYXFJYi9aQ1RTWmw3d09ZcHJIT3NianJaR21WN243QkVGU3lsUgplYzcvalRQWTFya3JLaEw2ZWJCeVBEWm1MckJQQWdNQkFBR2pXVEJYTUE0R0ExVWREd0VCL3dRRUF3SUNwREFQCkJnTlZIUk1CQWY4RUJUQURBUUgvTUIwR0ExVWREZ1FXQkJReUp2SWYremlBQ0xlYzZlUkxvbU5IdG5SUkNqQVYKQmdOVkhSRUVEakFNZ2dwcmRXSmxjbTVsZEdWek1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQ1o5c0NvT28wOQpVWnRpc2w2TTVLcmh4UFhKcFhmbXdSZ0wxd3doSTVFZWgvNy9pTnBkZlEzVW11Q2RiYWJ2YUUyQnB0bnBzN0dEClM2bytZNGo5dGRpSS9jUEZ0ekhXMVNCWU95SFEwYWptMTlTbkJtempUZzZKRDNTTVJiVVEzT1Q3VUtJVWpTQ0UKSnBRbWJoU2N6aHJKVjN1QnFLR1hzRk1xTDZRcXVaUlI5SVFJRmhzcnhidWp5cnNiZXZJbkhoQUFyWkp3ZnlVQgpTbmVjVkVsMm5NWFF3MjlRMDZIRVJkMWxJcW9KMFpoNzN2MWhmUDAyYmxma0xaL2tIaVE2ZGhLWDQrc0hVams0CnNlaVg2N3JnT2ZrRytBUFNGMEhaVEZqeUpHd3h3ZDFlem9DakdZbTRZc2oxNjUrSXQxZGYzWjN3QStoRFlCWmMKQjVPdDljejh2K0JHCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://192.168.0.104:6443
  name: kubernetes
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCVENDQWUyZ0F3SUJBZ0lJVWt4L2ZkQnhLbU13RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TkRFeE1EWXdNakEzTlRWYUZ3MHpOREV4TURRd01qRXlOVFZhTUJVeApFekFSQmdOVkJBTVRDbXQxWW1WeWJtVjBaWE13Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLCkFvSUJBUURZaDA4RFF3S2lhTWtMUHBNbElHQXArS21KVERLYnBLa2Y3NTNxS1VOeDdmRGVDU3lZWlpKK1VSRUoKUEsvKy9BMTFtdUw1cnZTSVplZDE4MC9xSGhUZGxYL3dQSlVMeG1pN1hNWHU2TG01SW95a1ZzTUJlTXhNUUVnSQovU01vdnZhdTZZY3Rhd2JCSi9ZME11MDkrNS9rTVAvTXdDVG9seGlETkJnTzZLZVJidEYyTWJLc0VxWTBrc3A5CkpnNm0xdUxUZEk2YTlTZ00yVUc1YUJZSzR5MTJ0WWFtY2pmVVZiSGxkR3lQNU9YeDY4RDNFNnI5NHhVUFVjR2kKc3BFekxMZkZQWnFmZGg3bDJjUEtnY3BiV1FPRklva0dKaGVJT1lPZ3pFOFZpbG0wS1dBbmhwWkFkQXZIbnVwRApNaUpHaSt1ejRTejJ3N2dJRzZPeXA1MkVyODVmQWdNQkFBR2pXVEJYTUE0R0ExVWREd0VCL3dRRUF3SUNwREFQCkJnTlZIUk1CQWY4RUJUQURBUUgvTUIwR0ExVWREZ1FXQkJSdGpIa0trL0I2bVRsNUc4ZFJMZEFReVN3QlRqQVYKQmdOVkhSRUVEakFNZ2dwcmRXSmxjbTVsZEdWek1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQ2R4WDE4VWhiTgpBcnE2UVg2K1B6WEFTS2FWWFBsWVNxV3VTS1dYUm1ZTjhvTnJZenpMZGZ0TGxmbWExWUZqaFpKaDY5Z2dTUlZ5CjlQTGU0OHRETDhFSDVMNFFOTDAwL2tZcThyaEtGei9lVC9TY1dITUMyUGZKaHM2Q3B5QzVtS3h5b0JGTE9rZlQKQ2hVOTlIalhkU0NmUmw1TXpGQWdSRVpQWWhhVHRRQWR5clZvcTdSUXZ2aFNWcVBBTmNobWJtMU5aVGIrVjhYSgo3SzFQUngrVmoraEFCRXhuSGR6d1BsYUdUVTJ0cjJUSTNGRjhkejZPdTVuV3E3aEdiNUl5Vk9RZ09QdHh0MW9UCko4ZG5aU0RDNWxjVnZtRzVQdElHcFFtYVkwMWtzMC8xejQ5ZXhEbllwb3dYMlhXZWk2eW5MU08wOEpjUHJKSHYKTGY2MlpQQUNoOU1WCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://127.0.0.1:59623
  name: kind-kind
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
- context:
    cluster: kind-kind
    user: kind-kind
  name: kind-kind
current-context: kind-kind
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lJR3VmQ1haOEJYNzR3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TkRBNU16QXdPREEyTWpkYUZ3MHlOVEE1TXpBd09ERXhNekZhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXlFdmFXMURaTkpHUE01ZkMKb1E4Wnplb0I2Nm5DQ3J3VndpM3I4blJwZG5WY0xBNGdNWVlXZnVtWm5IMEgxWEplUmdtanhGYUVDUDM1aEtQegphLzZMSDkvTnkrK2NJZXFGc1Q3NVZ2bGc3NWZVQ0Fvd0VoU3JQSHVZSDllTUVKRlV1d2xkcmZUdGxQaVFoM0MwCnhxd2lORkVSbXRmTnIrWjJTeXFLczAydmFtOWpuZW5Jdk90MzJMdWFmVWxwUHdGaGR0bkwrdTI3WXJVcTVSZDEKazhxSTBuS0NJdTVQeGxORUlyQnpEN29SbldRZEt1U0tyZFQyemI0dHdjRWlDckpLRzMwQVRZbjFOcGF5RW03egpZVkdTbjRhOFRCV1AzSUxpWDZhYjkwbEVWTGU3b3J2dmVrbTFlbzA0eDZoRWJETHlkaTg4amFsZnRjcjBqK1hlCkw1RS9hUUlEQVFBQm8xWXdWREFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JReUp2SWYremlBQ0xlYzZlUkxvbU5IdG5SUgpDakFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBdkZBcElheW9TT0I4Y0oyMWp2TE1pMldmY20zRUtteG1GajFZCjYwV045TXlHTlRwM2dEbVpUWmlKRGN3TXYvNlNZUG5KalMwTVpMb3dmQkZoOXhGM3BnUEJKMEF2UkNkcGl0aFoKelVYYlZZQlkzK1p0bHBwenU3aU9WVllkbkg0ZW9tTTJJS2VDUHpTamNScloxbk52V3QrYTk3ejU2T0hOcDdVcgpkZEFXN1pRQ21LT1ZGTzVvWlk3QkdmVTNpTzJvN1o4OURvT01RSzZjMElKSHBRNkxqQTNrdHA4YVBFKytsZ2IwCjNEcjEzRmVGUHQyNmYvL1Zwd1laMjlGN0JIeUFUK2F1SWphTy9MTGdLVUt4STdrQkFBRXZuempQeXp2QnlHbzkKcE9HU0ZuRE5yTVlOWmdpWWZnbTBrc3dLaWpIdmxQaXNFanhYaUUrWFpnQzRFOXA0WkE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBeUV2YVcxRFpOSkdQTTVmQ29ROFp6ZW9CNjZuQ0Nyd1Z3aTNyOG5ScGRuVmNMQTRnCk1ZWVdmdW1abkgwSDFYSmVSZ21qeEZhRUNQMzVoS1B6YS82TEg5L055KytjSWVxRnNUNzVWdmxnNzVmVUNBb3cKRWhTclBIdVlIOWVNRUpGVXV3bGRyZlR0bFBpUWgzQzB4cXdpTkZFUm10Zk5yK1oyU3lxS3MwMnZhbTlqbmVuSQp2T3QzMkx1YWZVbHBQd0ZoZHRuTCt1MjdZclVxNVJkMWs4cUkwbktDSXU1UHhsTkVJckJ6RDdvUm5XUWRLdVNLCnJkVDJ6YjR0d2NFaUNySktHMzBBVFluMU5wYXlFbTd6WVZHU240YThUQldQM0lMaVg2YWI5MGxFVkxlN29ydnYKZWttMWVvMDR4NmhFYkRMeWRpODhqYWxmdGNyMGorWGVMNUUvYVFJREFRQUJBb0lCQVFDcm9ETnVZNHg2ZXU5VgpxZ2hmc1d6UEFHQzg2aTBXdmF0M1E1b1ZtcUp6bW9Sc1MzNVNjUzc5ZUhUam5rOEVHb2VsUThWTUMwWC8zbi9iCnBCQ0V6UXV4T0RoRE13RjZIbGFJVmdtWStQNlN6bW9rcVhZZlNBNmlPTlZWRTRFMUFSSzFZWVVmOWV0TjV0OFEKN3dZMzVtODRuTzZVMjYybnQ3Wk5HaHJYSVEzYUNDb3pVNktwRkNjUGJtMWZaY0ZseDBzMkFpTCttVHhCQ3BJdQo5R2pJUS9sLy84QUNPcWx5Wi91VUFEa2tPQ1FtQTNvNEVVeEdaaGx5WXpJYkRpOEhXYWtLY3Jkc0p4T05MMk9GCkVYbTlrdXB4aUR1bmVLMHdLU0oxY1U5QWhjNGxqakRjWDcwN1lJaTVtNFhEc3FxYmt6MGoxVEl4TG1maDVhVzgKQTNxRlhZQnhBb0dCQVBwVlA5RjFMNStPU094N01rQjZLMHluTjNvUUhPaDM0a05BVU1LUTJ4YWlXZGs0TXdGago4WmMvSWdra004ZHA0YWhLM1JsS3lFL1J3YnJNYW5BSkR4QlZFeVgydi95UUFGNDhsbEM3UjU5T296aG10ZU0zCnR3Y0hvSHRWaUhUR29ZeXdGOE9pbk9udVJFUm1semJYT3pENDBoYWFKaCtNNnJCTFJBMDhhSHd6QW9HQkFNelUKb0hUY1h6cW0zNzAxTUUwMFVrWlpjb281SzR3L0lSVGIrc3d3Mk8razYvZE13QnA3bkxZQnBSTVZWTlhIbU5rVAorNGl4U04xYnpzU3gra0RIWDdmZWQybEFUK2g1cEpvcVV2WHAzcWtpWUtOSGxlSkFKYkEwM2pudS9kQUlORkNQCnNpWnVST1JmeW1nNnFYTXg0NXBjbnh6cUFuUGFzaXVrdE5MY1ZEbnpBb0dCQU8yOVNINkQ3Rlo3cW9YcitpMkIKMk4xVGNUeGJVUmoxd2N4Y3FGWWZlL0ppL1RGdVRnSmtDR3k3YUhlR0NpYTRSN2FzWW81Q2x6bzIydVdzZk9rcApzVVN4aHgzbTJTM2pGSFpxMDlhWUJjMGx3WjB1N2s1Nyt6YVI1N2M1NC80REppbVdrdnNZMUN6V083ODZMeUhHCkJsRGIvYW01ZTdzNitTZTBVMHkrc2Z4QkFvR0JBSU16czNBSGRLeENGZERCa0NYejNMdVpNZ2dkNUtvYUNkdXQKcUxGQW5NU3NORVdkRVBRbHQ5VFJxdVpWWkpqbkdCMzhjY00ySkFFK2ZHeDd3RnZjR1pEU1hGUzcwRE9PTDRSYwpsZlZWRDczdytrdThYK0tqeWtCYkxQbVkvMVZRM0FtNmNaZXlURWlvbnlNeWFEWVVmOER4a1MzWkt5Y0FyOTNLCnk5VEJNdVpIQW9HQVdYc1VGK25QbWdHc1l5bHpUSzl6NStxNmloelI1enhvNUVCd2ZHTlpBRXBSMVphSHdVcXIKWGVzOFBSWWp0c3FORTBsdWJNOGs1QS9hM2Q4QzhBdU5IdjBXSFRTQUVlRUlqeWkxeTVoakJyaklhWXdPNGZmTgpOK0I3Rm9OOVJQbDdUWHRkckNmck1hbklYU3JLOXdKMlpXRm1LaWNJT1lVUzBZYlNaOHh2emhBPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
- name: kind-kind
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURLVENDQWhHZ0F3SUJBZ0lJRHNFcTVnRyt1VkV3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TkRFeE1EWXdNakEzTlRWYUZ3MHlOVEV4TURZd01qRXlOVFZhTUR3eApIekFkQmdOVkJBb1RGbXQxWW1WaFpHMDZZMngxYzNSbGNpMWhaRzFwYm5NeEdUQVhCZ05WQkFNVEVHdDFZbVZ5CmJtVjBaWE10WVdSdGFXNHdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFERDdMQTIKdUVDcE5meDBDdGE3RmFHWTF1YjF4cTdoRGdGTVhvTHk2NkxKbXBKT3lYUllHUFU1LzJHUlRId1UvRnVzSWYvQgphN2RzRjhWL0MvZStlKzZUR3BLOEZBR3ZsTmp4dkpEaENOVlRSTTR0bXpSR2lTTXo5cTJtMURWcW5GMDNDTjBzCjhXZXJIWmdjNVIyQ3A2WkYyMUdXTDIxVWU0UENvN2k5alZkNXJQNERHTWRCTkFYSVlGaXg1VWE2akd5NkpncTgKWXNDTms1Wko4Z0tIOFFIa0VtckZzWmV1NEI4MmxGUkdpWHN6Sjc2VVdvSE1TNWlJWGtSNWlrUmNJVEFvUFVxRgoyZEtiSnZxODN5RFdQVzVrVDNoMmh1Ri9ZWkppQ0xwVDJlYmRQSUxjR1lSczd0dlhlbXpBdXpVaU8xcUlBeitsClAvaSsyQmNNRDFQY0FQaUpBZ01CQUFHalZqQlVNQTRHQTFVZER3RUIvd1FFQXdJRm9EQVRCZ05WSFNVRUREQUsKQmdnckJnRUZCUWNEQWpBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkcyTWVRcVQ4SHFaT1hrYgp4MUV0MEJESkxBRk9NQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUURHWHJuUWduN2I4elJlYlZ4clpGTndpWktPCmFYaDZKTVNIbm0wNm9iYThkMDR2YVVCbzZwT01lOE9PNTcvclRNcWR4cXhFNTAxMlY0eXZBNng4RzgzY01KdDgKaGNkQ2ZlcjFQbE9YZ1RaVlBnV2hiVHJJSjc2M3hUQm1IOEJYT3hsQ1p4ZXAzS2xZWk9nSGxpeFJtRnYwZklMQgpwdERCdGE4R3BUcnRqaFZ5eWxHMjA4ajJwK3oxOE5Sc3A2bHdmRVZnenRFbTBFeFM2MFA5VjY1ZzhoR2M1UVpoCm5RSVhTSDVBZGxVYTYrajRuSndPOG01Rmk4NGFMbHFzR3czalJCeHZnUkxYUk81clBXWHpJL1NrL3ZpcytlTDEKbTVnMFVmQXJLQU4wc096bTFoTGR4YVZ5QUZUVU5XS204aTJCd0U3SXk2bXJ6VnZMVWJyVjNpR2s3aEZ6Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdyt5d05yaEFxVFg4ZEFyV3V4V2htTmJtOWNhdTRRNEJURjZDOHV1aXlacVNUc2wwCldCajFPZjloa1V4OEZQeGJyQ0gvd1d1M2JCZkZmd3Yzdm52dWt4cVN2QlFCcjVUWThieVE0UWpWVTBUT0xaczAKUm9rak0vYXRwdFExYXB4ZE53amRMUEZucXgyWUhPVWRncWVtUmR0UmxpOXRWSHVEd3FPNHZZMVhlYXorQXhqSApRVFFGeUdCWXNlVkd1b3hzdWlZS3ZHTEFqWk9XU2ZJQ2gvRUI1QkpxeGJHWHJ1QWZOcFJVUm9sN015ZStsRnFCCnpFdVlpRjVFZVlwRVhDRXdLRDFLaGRuU215YjZ2TjhnMWoxdVpFOTRkb2JoZjJHU1lnaTZVOW5tM1R5QzNCbUUKYk83YjEzcHN3THMxSWp0YWlBTS9wVC80dnRnWERBOVQzQUQ0aVFJREFRQUJBb0lCQVFDYzdxTWUwV3NKbm5KKwpTSWhEQmtxUDcrTERqc2RaQVN6TkRRNzZvUCtkV0RCRTUxeEhqSVl3Vkh6RU0yMVlLZU1MOTVleVNDTjljM1VBCkZJZjJqYkpGSmczT2xIL2RNZTZyZ296Ums0Kzd5T3NVNExKNHBUUUxWVlUyd2RlZmMydSt2MXpadU90K3hvK20KNVdaRDF5RjU1dmhzd2NSaTNTUm03VmoyaTVZN29JeGhvVjl6S2dXc2VBbllkeXgzU3Z5WTBXaGZ6eENuSTUxQgovVUVTd1hDak5LVUdXUDZSNElDTDYwTGw1Z1RoOWtrMTh1S1BrWWJLSDlmRGw2MjNBTC9tSkZpZ3g1NUxJL25NCkhORUpBMjdYcnhUTCtVK2M3TXdvd0tKM2o0V210ZkZOdDFZeHhjS3l2OXB0R1FPRUVTY25oRG9zNDNkMDdaMGIKaWovQ1RIN2hBb0dCQU9UWE9EUWlrYUsrcXV3NVdQZTU2djZXWk1QeGdtQzFCbThOOGJNQm9aT2xVS3pUbS9nSgpuSFVaMWYyQ1V4UmYxTysycUJwa0RmRmw4MmZVVTBzNk0zZXc4WHNyZXdFK01MYmUrcjc3UWFaVVAvQjAyYW1UCkRwazQrakl6K0lWSGZIbXVTWkIvd08zZTNHQ1ZPeUIxb2p6SnBJYXRTdWdXU243M3ljVDFXOW10QW9HQkFOc3QKWStaMTdrRldVdVJ5T1F4WVJMYkxYbWpIQ1V3OWlhdFk3L1ZtL1VLcDB3MmxlSDJHQ2NRRCtIYzliS3FTeDJWSQpoRWIwZ1dPZldNeUpES0w2QVZtVzR2SkdXR0tSeVJ5Unl0ZU5PZk56V2dReFVtcFkrZDZzRFZYN0dacXJ5UWRNClZtODVpaTBENk1OcCtXUkpadHczWlBkZHNwOWhIMzhWeVdGUGVXM05Bb0dBR0gzSU1CdzdCZlh1Q1JZaUpYRXEKYTFEaE8rOVBDdGFVOTdIQVdtNGtRczhBa1Y2Y1pMRnlvejIrbjBFaGJ4N0toVlZCTElIazFCOGJLOU9Yam9lTwpGcE5EWlBGRVd0K3pDdjlXU3JaTlVtWFY2Z0EzZzJTUHZXcFJyS25QUVVSaldBcUZLUWZqT0JJUDkrNUF3N3FUClFIbzhONFc0YkpwbUlxeVdWWlFFM29rQ2dZQTNlTzUrNXJ3dGh6YWxvUTgxUTZYb2llSlVMSVA2NnR4TUpNOWUKMGZrcGhTZm9uVWU0cFZNVmJGZlhmaEZodnBKKzNQSzFycTZNMDBpN1E3aVNDeXFLVFRrVlRwNlNIQW5GbEZTOQpaMzRTVXRDbW5RRVo3M2tXVlg5dWtvWHhjcWNIbE5lUGdRV3F6UUY5YS9YMTN1b010R3gyZXgxNVh6Q0VqclFRClQvZ1F4UUtCZ1FES1FkSzhFalFKc3lWTUJ1ZmlJdGdaalVZK0h3YUFoMC9NeHNvZmxHYWJiUkNnckdFaUpwRmkKYVc0Y1M2dUV1SVlDcStZRVo5b3FJOHd2WUlMOGU0bGwwV1BCTDQwL2pqZFkrSmVzYmNmbmIxMW5vVCtOUXdqSQpYdWh1VER4WDBuVjJXNGE4Ni9vaUtQQnh1ZmFwSEhrWGRTckJjYmcyODFQRTJZd1VLYVQzWVE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
`,
			ActionTimeoutSeconds: 30,
		},
	}

	for _, cluster := range clusters {
		// 插入或更新 K8sCluster
		if err := m.db.Where("name = ?", cluster.Name).FirstOrCreate(&cluster).Error; err != nil {
			log.Printf("populateMockData: 插入 K8sCluster 失败: %v\n", err)
			continue
		}

		log.Printf("populateMockData: 初始化 Kubernetes 集群成功，ClusterID: %d\n", cluster.ID)
	}
}
