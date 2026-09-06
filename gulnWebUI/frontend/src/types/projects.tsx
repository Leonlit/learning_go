export interface NewProjectResponse {
    uuid: string;
};

export interface Project {
	uuid: string
	name: string
	created_at: string
	updated_at: string
}

export interface ProjectCount {
    count: number;
};