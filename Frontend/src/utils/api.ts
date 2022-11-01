import { useQuery, useQueryClient } from '@tanstack/react-query';
import { Containers, SystemInfo } from '../types/types';

export function GetSystemInfo() {
    const queryClient = useQueryClient();
    
    const systemInfoQuery = useQuery<SystemInfo>(
      ["systemInfo"], 
      () => fetch("http://localhost:8080/api/system-stats").then(response => response.json()),
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
      () => fetch("http://localhost:8080/api/containers").then(response => (response.status === 200) ? response.json(): (function(){throw new Error("Error")})()),
      {
        staleTime: 60000,
        refetchInterval: 60000,
        refetchOnWindowFocus: true
      }
    );

    return containersQuery;
}