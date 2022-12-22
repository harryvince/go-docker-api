import { useQuery, useQueryClient, useMutation } from '@tanstack/react-query';
import { Containers, SystemInfo } from '../types/types';

const API_URL = import.meta.env.VITE_API_URL;

export function GetSystemInfo() {
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

export function CreateContainer(type: "mysql" | "ubuntu" | "postgres"){
    const queryClient = useQueryClient();
    const createMutation = useMutation<Containers>(
      [`${type}`], 
      () => fetch(`${API_URL}/${type}`, 
      {
        method: 'POST',
        body: JSON.stringify({port: Math.floor(1000 + Math.random() * 9000).toString()})
      }
      ).then(response => (response.status === 201) ? response.json(): (function(){throw new Error("Error")})()),
      {
        onSuccess: () => queryClient.invalidateQueries(["containers", "system-info"]),
      }
    );

    return createMutation;
}