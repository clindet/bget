package types

type GdcEndpoints struct {
	Status      bool
	Projects    bool
	Cases       bool
	Files       bool
	Annotations bool
	Data        bool
	Manifest    bool
	Slicing     bool
	Legacy      bool
	ExtraParams GdcExtraParams
}

type GdcExtraParams struct {
	RemoteName bool
	Token      string
	Query      string
	Size       int
	From       int
	Sort       string
	Filter     string
	Pretty     bool
	JSON       bool
	Format     string
	Fields     string
	Expand     string
	Facets     string
}

type GdcStatus struct {
	Commit      string `json:"commit"`
	DataRelease string `json:"data_release"`
	Status      string `json:"status"`
	Tag         string `json:"tag"`
	Version     int    `json:"version"`
}

type GdcProjects struct {
	Data struct {
		Hits []struct {
			DbgapAccessionNumber string   `json:"dbgap_accession_number"`
			DiseaseType          []string `json:"disease_type"`
			ID                   string   `json:"id"`
			Name                 string   `json:"name"`
			PrimarySite          []string `json:"primary_site"`
			ProjectID            string   `json:"project_id"`
			Releasable           bool     `json:"releasable"`
			Released             bool     `json:"released"`
			State                string   `json:"state"`
		} `json:"hits"`
		Pagination struct {
			Count int    `json:"count"`
			From  int    `json:"from"`
			Page  int    `json:"page"`
			Pages int    `json:"pages"`
			Size  int    `json:"size"`
			Sort  string `json:"sort"`
			Total int    `json:"total"`
		} `json:"pagination"`
	} `json:"data"`
	Warnings struct{} `json:"warnings"`
}

type GdcCases struct {
	Data struct {
		Hits []struct {
			AliquotIds            []string    `json:"aliquot_ids"`
			CaseID                string      `json:"case_id"`
			CreatedDatetime       string      `json:"created_datetime"`
			DaysToLostToFollowup  interface{} `json:"days_to_lost_to_followup"`
			DiagnosisIds          []string    `json:"diagnosis_ids"`
			DiseaseType           string      `json:"disease_type"`
			ID                    string      `json:"id"`
			IndexDate             string      `json:"index_date"`
			LostToFollowup        interface{} `json:"lost_to_followup"`
			PrimarySite           string      `json:"primary_site"`
			SampleIds             []string    `json:"sample_ids"`
			State                 string      `json:"state"`
			SubmitterAliquotIds   []string    `json:"submitter_aliquot_ids"`
			SubmitterDiagnosisIds []string    `json:"submitter_diagnosis_ids"`
			SubmitterID           string      `json:"submitter_id"`
			SubmitterSampleIds    []string    `json:"submitter_sample_ids"`
			UpdatedDatetime       string      `json:"updated_datetime"`
		} `json:"hits"`
		Pagination struct {
			Count int    `json:"count"`
			From  int    `json:"from"`
			Page  int    `json:"page"`
			Pages int    `json:"pages"`
			Size  int    `json:"size"`
			Sort  string `json:"sort"`
			Total int    `json:"total"`
		} `json:"pagination"`
	} `json:"data"`
	Warnings struct{} `json:"warnings"`
}

type GdcFiles struct {
	Data struct {
		Hits []struct {
			Access          string      `json:"access"`
			Acl             []string    `json:"acl"`
			CreatedDatetime string      `json:"created_datetime"`
			DataCategory    string      `json:"data_category"`
			DataFormat      string      `json:"data_format"`
			DataRelease     string      `json:"data_release"`
			DataType        string      `json:"data_type"`
			ErrorType       interface{} `json:"error_type"`
			FileID          string      `json:"file_id"`
			FileName        string      `json:"file_name"`
			FileSize        int         `json:"file_size"`
			ID              string      `json:"id"`
			Md5sum          string      `json:"md5sum"`
			State           string      `json:"state"`
			StateComment    interface{} `json:"state_comment"`
			SubmitterID     string      `json:"submitter_id"`
			Type            string      `json:"type"`
			UpdatedDatetime string      `json:"updated_datetime"`
			Version         string      `json:"version"`
		} `json:"hits"`
		Pagination struct {
			Count int    `json:"count"`
			From  int    `json:"from"`
			Page  int    `json:"page"`
			Pages int    `json:"pages"`
			Size  int    `json:"size"`
			Sort  string `json:"sort"`
			Total int    `json:"total"`
		} `json:"pagination"`
	} `json:"data"`
	Warnings struct{} `json:"warnings"`
}

type GdcAnnotations struct {
	Data struct {
		Hits []struct {
			AnnotationID          string      `json:"annotation_id"`
			CaseID                string      `json:"case_id"`
			CaseSubmitterID       string      `json:"case_submitter_id"`
			Category              string      `json:"category"`
			Classification        string      `json:"classification"`
			CreatedDatetime       string      `json:"created_datetime"`
			EntityID              string      `json:"entity_id"`
			EntitySubmitterID     string      `json:"entity_submitter_id"`
			EntityType            string      `json:"entity_type"`
			ID                    string      `json:"id"`
			LegacyCreatedDatetime string      `json:"legacy_created_datetime"`
			LegacyUpdatedDatetime interface{} `json:"legacy_updated_datetime"`
			Notes                 string      `json:"notes"`
			State                 string      `json:"state"`
			Status                string      `json:"status"`
			SubmitterID           string      `json:"submitter_id"`
			UpdatedDatetime       string      `json:"updated_datetime"`
		} `json:"hits"`
		Pagination struct {
			Count int    `json:"count"`
			From  int    `json:"from"`
			Page  int    `json:"page"`
			Pages int    `json:"pages"`
			Size  int    `json:"size"`
			Sort  string `json:"sort"`
			Total int    `json:"total"`
		} `json:"pagination"`
	} `json:"data"`
	Warnings struct{} `json:"warnings"`
}
