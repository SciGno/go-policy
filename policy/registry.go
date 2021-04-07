package policy

// Validator interface validates statement data against request data.
type Validator interface {
	Validate(stmntData, requestData interface{}) bool
}
type ValidatorMap struct {
	validators map[string]Validator
}

func NewValidatorMap(validatorMap map[string]Validator) ValidatorMap {
	return ValidatorMap{
		validators: validatorMap,
	}
}

func (vm *ValidatorMap) AddValidator(name string, validator Validator) {
	vm.validators[name] = validator
}

func (vm *ValidatorMap) RemoveValidator(name string) {
	delete(vm.validators, name)
}

func (vm *ValidatorMap) GetValidator(name string) Validator {
	return vm.validators[name]
}

func (vm *ValidatorMap) GetValidators() map[string]Validator {
	return vm.validators
}

type Registry struct {
	maps map[string]ValidatorMap
}

// NewRegistry return an empty Registry if a ValidatorMap with "resource" and "action"
// are not provided in the validatorMap parameter.
func NewRegistry(validatorMap map[string]ValidatorMap) Registry {

	registry := Registry{
		maps: map[string]ValidatorMap{},
	}

	if validatorMap != nil {
		registry.maps = validatorMap
	}

	if _, ok := registry.maps[string(RESOURCE)]; !ok {
		registry.AddValidatorMap(
			string(RESOURCE),
			NewValidatorMap(map[string]Validator{"default": &DelimitedValidator{}}),
		)
	}

	if _, ok := registry.maps[string(ACTION)]; !ok {
		registry.AddValidatorMap(
			string(ACTION),
			NewValidatorMap(map[string]Validator{"default": &ActionValidator{}}),
		)
	}

	return registry
}

func (r *Registry) AddValidatorMap(metaName string, validatorMap ValidatorMap) {
	r.maps[metaName] = validatorMap
}

func (r *Registry) GetValidatorMap(metaName string) ValidatorMap {
	return r.maps[metaName]
}

func (r *Registry) RemoveValidatorMap(metaName string) ValidatorMap {
	validatorMap := r.maps[metaName]
	delete(r.maps, metaName)
	return validatorMap
}

func (r *Registry) AddValidator(metaName, validatorName string, validator Validator) {
	r.maps[metaName] = NewValidatorMap(map[string]Validator{validatorName: validator})
}

func (r *Registry) GetValidator(metaName, validatorName string) Validator {
	validatorMap := r.maps[metaName]
	return validatorMap.GetValidator(validatorName)
}

func (r *Registry) RemoveValidator(metaName, validatorName string) Validator {
	validatorMap := r.maps[metaName]
	validator := validatorMap.GetValidator(validatorName)
	delete(validatorMap.validators, validatorName)
	return validator
}

func (r *Registry) MetaNameExists(metaName string) bool {
	if len(r.GetValidatorMap(metaName).validators) == 0 {
		return false
	}
	return true
}
