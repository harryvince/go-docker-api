import { useQuery, useQueryClient } from '@tanstack/react-query';
import { SystemInfo } from './types/types';

export default function App() {
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



  return (
    <div className="hero min-h-screen" style={{ backgroundColor: '#202A44' }}>
      <div className="hero-overlay bg-opacity-60"></div>
      <div className="hero-content text-center text-neutral-content">
        <div className="max-w-md">
          <h1 className="mb-5 text-5xl font-bold">System Information</h1>
          <p className="mb-5">{systemInfoQuery.isFetching ? (
            <span>Refreshing...</span>
          ): (
            <>
              <span>CPU Free: {Math.round(Number(systemInfoQuery.data?.CPU.Free))}%</span>
              <div className="pt-2"></div>
              <span>CPU Used: {Math.round(Number(systemInfoQuery.data?.CPU.Used))}%</span>
              <div className="pt-2"></div>
              <span>Memory Free: {Math.round(Number(systemInfoQuery.data?.Memory.Free))}%</span>
              <div className="pt-2"></div>
              <span>Memory Used: {Math.round(Number(systemInfoQuery.data?.Memory.Used))}%</span>
            </>
          )}</p>
        </div>
      </div>
    </div>
  )
}