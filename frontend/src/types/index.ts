export interface AboutInfo {
  version: string;
}

export interface AnalysisFormData {
  files: File[];
  startDate?: string;
  startTime?: string;
  endDate?: string;
  endTime?: string;
  limit?: number;
  window?: number;
}
