export interface AssessmentMinimal {
	uuid: string
	name: string
	start_datetime: string
	updated_datetime: string
    end_datetime: string
}

export interface Assessment {
	uuid: string
	name: string
	start_datetime: string
	updated_datetime: string
    end_datetime: string
    type: string
    phase: string
}