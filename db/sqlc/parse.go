package db

func ParseDoctorToUpdateParams(doctor Doctor) UpdateDoctorParams {
	return UpdateDoctorParams{
		ID:          doctor.ID,
		Name:        doctor.Name,
		Crm:         doctor.Crm,
		AverageRate: doctor.AverageRate,
	}

}
