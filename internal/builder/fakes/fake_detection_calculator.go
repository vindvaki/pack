package fakes

import (
	"github.com/buildpacks/pack/builder"
	"github.com/buildpacks/pack/pkg/dist"
)

type FakeDetectionCalculator struct {
	ReturnForOrder builder.DetectionOrder

	ErrorForOrder error

	ReceivedTopOrder dist.Order
	ReceivedLayers   dist.BuildpackLayers
	ReceivedDepth    int
}

func (c *FakeDetectionCalculator) Order(
	topOrder dist.Order,
	layers dist.BuildpackLayers,
	depth int,
) (builder.DetectionOrder, error) {
	c.ReceivedTopOrder = topOrder
	c.ReceivedLayers = layers
	c.ReceivedDepth = depth

	return c.ReturnForOrder, c.ErrorForOrder
}
