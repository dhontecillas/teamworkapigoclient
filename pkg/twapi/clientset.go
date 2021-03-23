package twapi

import (
	"context"

	"github.com/dhontecillas/teamworkapigoclient/pkg/openapi/pmv1"
	"github.com/dhontecillas/teamworkapigoclient/pkg/openapi/projv1"
	"github.com/dhontecillas/teamworkapigoclient/pkg/openapi/projv2"
	"github.com/dhontecillas/teamworkapigoclient/pkg/openapi/projv3"
)

const (
	TIMEFORMATSTR    = "2006-01-02T15:04:05Z07:00"
	V1_DATEFORMATSTR = "20060102"
	V1_TIMEFORMATSTR = "15:04"
)

type ClientSet struct {
	Pm    *pmv1.APIClient
	PmCtx context.Context
	V1    *projv1.APIClient
	V1Ctx context.Context
	V2    *projv2.APIClient
	V2Ctx context.Context
	V3    *projv3.APIClient
	V3Ctx context.Context
}

func NewClientSet(scheme string, host string, apiKey string, ctx context.Context) *ClientSet {
	cs := ClientSet{}
	confV3 := projv3.NewConfiguration()
	/* this should be done this way, but is not taken into account:
	conf.Scheme = "https"
	conf.Host = "davidhontecillastest.teamwork.com"
	*/
	confV3.Scheme = scheme
	confV3.Host = host
	cs.V3Ctx = context.WithValue(ctx, projv3.ContextBasicAuth,
		projv3.BasicAuth{
			UserName: apiKey,
			Password: "XXX",
		})
	cs.V3 = projv3.NewAPIClient(confV3)

	confV2 := projv2.NewConfiguration()
	confV2.Scheme = scheme
	confV2.Host = host
	cs.V2Ctx = context.WithValue(ctx, projv2.ContextBasicAuth,
		projv2.BasicAuth{
			UserName: apiKey,
			Password: "XXX",
		})
	cs.V2 = projv2.NewAPIClient(confV2)

	confV1 := projv1.NewConfiguration()
	confV1.Scheme = scheme
	confV1.Host = host
	cs.V1Ctx = context.WithValue(ctx, projv1.ContextBasicAuth,
		projv1.BasicAuth{
			UserName: apiKey,
			Password: "XXX",
		})
	cs.V1 = projv1.NewAPIClient(confV1)

	confVpm := pmv1.NewConfiguration()
	confVpm.Scheme = scheme
	confVpm.Host = host
	cs.PmCtx = context.WithValue(ctx, pmv1.ContextBasicAuth,
		pmv1.BasicAuth{
			UserName: apiKey,
			Password: "XXX",
		})
	cs.Pm = pmv1.NewAPIClient(confVpm)

	return &cs
}
