import { useQuery } from '@tanstack/react-query';
import { aboutApi } from '../api/client';

export const useAbout = () => {
  return useQuery({
    queryKey: ['about'],
    queryFn: aboutApi.getAbout,
  });
};
