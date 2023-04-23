package client

type CreateResourceRequest struct {
	IntegrationType string `json:"integrationType"`
	Name            string `json:"name"`
	Configuration   string `json:"config"`
}

type UpdateResourceRequest struct {
	IntegrationType string `json:"integrationType"`
	Configuration   string `json:"config"`
}

type CreateResourceResponse struct {
	ID string `json:"id"`
}

type UpdateResourceResponse struct {
	ID string `json:"id"`
}

const (
	PostgresIntegrationType = "postgres"
)

// sessions
type CreateSessionRequest struct {
	SessionName       string `json:"sessionName"`
	ResourceName      string `json:"resourceName"`
	ClusterName       string `json:"clusterName"`
	AuthorizationName string `json:"authorizationName"`
	SessionTTL        string `json:"sessionTTL"`
	SessionType       string `json:"sessionType"`
}

type CreateSessionResponse struct {
	ID string `json:"id"`
}

type UpdateSessionRequest = CreateSessionRequest

// type UpdateSessionRequest struct {
// 	SessionName       string `json:"sessionName"`
// 	ClusterName       string `json:"clusterName"`
// 	AuthorizationName string `json:"authorizationName"`
// 	SessionTTL        string `json:"sessionTTL"`
// }

// type UpdateSessionRequest struct {
// 	*CreateSessionRequest
// }

type UpdateSessionResponse struct {
	ID string `json:"id"`
}
