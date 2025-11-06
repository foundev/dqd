import axios from 'axios';
import type { AboutInfo } from '../types';

const API_BASE_URL = import.meta.env.VITE_API_URL || '/api';

export const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const aboutApi = {
  getAbout: async (): Promise<AboutInfo> => {
    const response = await apiClient.get<AboutInfo>('/about.json');
    return response.data;
  },
};

export const profileApi = {
  analyzeProfile: async (file: File): Promise<string> => {
    const formData = new FormData();
    formData.append('profile1', file);

    const response = await apiClient.post('/profile', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },

  analyzeSimpleProfile: async (file: File): Promise<string> => {
    const formData = new FormData();
    formData.append('profile1', file);

    const response = await apiClient.post('/simple-profile', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },

  compareProfiles: async (file1: File, file2: File): Promise<string> => {
    const formData = new FormData();
    formData.append('compare_profile', file1);
    formData.append('compare_profile', file2);

    const response = await apiClient.post('/profiles', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },
};

export const queriesApi = {
  analyzeQueries: async (
    file: File,
    options: {
      startDate?: string;
      startTime?: string;
      endDate?: string;
      endTime?: string;
      limit?: number;
      window?: number;
    }
  ): Promise<string> => {
    const formData = new FormData();
    formData.append('queriesjson', file);
    formData.append('start_date', options.startDate || '');
    formData.append('start_time', options.startTime || '00:00');
    formData.append('end_date', options.endDate || '');
    formData.append('end_time', options.endTime || '00:00');
    formData.append('limit', String(options.limit || 5));
    formData.append('window', String(options.window || 60000));
    formData.append('query_report_type', 'INTERACTIVE');

    const response = await apiClient.post('/queriesjson', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },
};

export const reproApi = {
  generateSchema: async (
    file: File,
    options: {
      records?: number;
      timeout?: number;
      defaultCtasFormat?: string;
      nasPath?: string;
      columnDefYaml?: string;
    }
  ): Promise<Blob> => {
    const formData = new FormData();
    formData.append('profile', file);
    formData.append('records', String(options.records || 20));
    formData.append('timeout', String(options.timeout || 60));
    if (options.defaultCtasFormat) {
      formData.append('defaultCtasFormat', options.defaultCtasFormat);
    }
    if (options.nasPath) {
      formData.append('nasPath', options.nasPath);
    }
    if (options.columnDefYaml) {
      formData.append('columnDefYaml', options.columnDefYaml);
    }

    const response = await apiClient.post('/reproduction', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      responseType: 'blob',
    });
    return response.data;
  },
};

export const systemApi = {
  analyzeIOStat: async (file: File): Promise<string> => {
    const formData = new FormData();
    formData.append('iostatfile', file);

    const response = await apiClient.post('/iostat', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },

  analyzeTop: async (file: File): Promise<string> => {
    const formData = new FormData();
    formData.append('ttop', file);

    const response = await apiClient.post('/ttop', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },
};
