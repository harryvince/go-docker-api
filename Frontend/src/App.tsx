import { useQuery, useQueryClient } from '@tanstack/react-query';
import ContainerOptions from './components/ContainerOptions';
import SystemUtilization from './components/SystemUtilization';
import { SystemInfo } from './types/types';
import Button from './components/Button';
import postgreSQL from "./assets/postgreSQL.png";
import mySQL from "./assets/mySQL.png";
import ubuntu from "./assets/ubuntu-logo.png";

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
    <div className="min-h-screen" style={{ backgroundColor: '#202A44' }}>
      <div className="grid grid-cols-3 gap-4">
        <div id="column1" className='flex justify-center'></div>
        <div id="column2" className="flex justify-center">
          <div className="pt-8">
            <ContainerOptions />
            <div className="mt-4">
              <SystemUtilization />
            </div>
            <div className='mt-4 flex justify-center'>
              <Button image={postgreSQL} container_type={"postgres"} />
              <Button image={mySQL} container_type={"mysql"} />
              <Button image={ubuntu} container_type={"ubuntu"} />
            </div>
          </div>
        </div>
        <div id="column3"></div>
      </div>
    </div>
  )
}