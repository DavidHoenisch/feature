package main

type CheckerImpl struct{}

func NewFeater() *CheckerImpl {

	return &CheckerImpl{}
}

func (c *CheckerImpl) CheckStateModel(v string) (*Flag, error) {
	return parseFlagToModel(v)
}

func (c *CheckerImpl) CheckStateInt(v string) (int, error) {
	return GetFlagFromEnv(v).ToInt()
}

func (c *CheckerImpl) CheckStateFloat(v string) (float64, error) {
	return GetFlagFromEnv(v).ToFloat64()
}

func (c *CheckerImpl) CheckStateBool(v string) (bool, error) {
	return GetFlagFromEnv(v).ToBool()
}

func (c *CheckerImpl) CheckStateString(v string) (string, error) {
	return v, nil
}
