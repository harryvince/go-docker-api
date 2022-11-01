import { useQuery, useQueryClient } from '@tanstack/react-query';
import { Containers, SystemInfo } from '../types/types';

const API_URL = import.meta.env.VITE_API_URL;

export function GetSystemInfo() {
    const queryClient = useQueryClient();
    
    const systemInfoQuery = useQuery<SystemInfo>(
      ["systemInfo"], 
      () => fetch(`${API_URL}/system-stats`).then(response => (response.status === 200) ? response.json(): (function(){throw new Error("Error")})()),
      {
        staleTime: 60000,
        refetchInterval: 60000,
        refetchOnWindowFocus: true
      }
    );

    return systemInfoQuery;
}

export function GetContainers() {
    const queryClient = useQueryClient();
    
    const containersQuery = useQuery<Containers>(
      ["containers"], 
      () => fetch(`${API_URL}/containers`).then(response => (response.status === 200) ? response.json(): (function(){throw new Error("Error")})()),
      {
        staleTime: 60000,
        refetchInterval: 60000,
        refetchOnWindowFocus: true
      }
    );

    return containersQuery;
}