import { GetSystemInfo } from "../utils/api"

export default function SystemUtilization() {
    const { data, error, isLoading } = GetSystemInfo();
    
    return (
        <div className="stats shadow">
  
          <div className="stat">
            <div className="stat-title">Total CPU Usage</div>
            <div className="stat-value text-primary">{isLoading ? ('Loading...'): `${Math.round(Number(data?.CPU.Used))}%`}</div>
            <div className="stat-desc">{isLoading ? ('Loading...'): `Free: ${Math.round(Number(data?.CPU.Free))}%`}</div>
          </div>

          <div className="stat">
            <div className="stat-title">Total Memory Usage</div>
            <div className="stat-value text-secondary">{isLoading ? ('Loading...'): `${Math.round(Number(data?.Memory.Used))}%`}</div>
            <div className="stat-desc">{isLoading ? ('Loading...'): `Free: ${Math.round(Number(data?.Memory.Free))}%`}</div>
          </div>

        </div>
    )
}