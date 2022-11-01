import { Container } from "../types/types";
import { GetContainers } from "../utils/api";
import dockerContainer from "../assets/dockerContainer.png";
import postgreSQL from "../assets/postgreSQL.png";
import mySQL from "../assets/mySQL.png";
import ubuntu from "../assets/ubuntu-logo.png";

export default function ContainerOptions() {
    const { data, isLoading, isError } = GetContainers();
    
    return (
        <>
        <p className="text-white text-2xl text-center font-bold pb-4 italic">Live Containers</p>
        <div className="overflow-x-auto w-full">
        <table className="table w-full">
          <thead>
            <tr>
              <th>Type</th>
              <th>Container</th>
              <th>Port</th>
            </tr>
          </thead>
          <tbody>
              {isLoading ? (
                <tr>
                    <td>
                      <div className="flex items-center space-x-3">
                        <div className="avatar">
                          <div className="mask w-12 h-12">
                            <img src={dockerContainer} alt="LoadingItemHolder" />
                          </div>
                        </div>
                      </div>
                    </td>
                    <td>Loading...</td>
                    <td>Loading...</td>
                </tr>
                ) : (
                    isError ? (
                        <tr>
                            <td>
                              <div className="flex items-center space-x-3">
                                <div className="avatar">
                                  <div className="mask w-12 h-12">
                                    <img src={dockerContainer} alt="LoadingItemHolder" />
                                  </div>
                                </div>
                              </div>
                            </td>
                            <td>No Containers Found</td>
                            <td>N/A</td>
                        </tr>
                    ) : (
                        data!.containers.map((container: Container) => (
                            <tr key={container.ID}>
                                {container.Name.split("-")[0] === "postgres" && (
                                    <td>
                                      <div className="flex items-center space-x-3">
                                        <div className="avatar">
                                          <div className="mask w-12 h-12">
                                            <img src={postgreSQL} alt="PostgreSQL" />
                                          </div>
                                        </div>
                                      </div>
                                    </td>
                                )}
                                {container.Name.split("-")[0] === "mysql" && (
                                    <td>
                                      <div className="flex items-center space-x-3">
                                        <div className="avatar">
                                          <div className="mask w-12 h-12">
                                            <img src={mySQL} alt="MySQL" />
                                          </div>
                                        </div>
                                      </div>
                                    </td>
                                )}
                                {container.Name.split("-")[0] === "ubuntu" && (
                                    <td>
                                      <div className="flex items-center space-x-3">
                                        <div className="avatar">
                                          <div className="mask w-12 h-12">
                                            <img src={ubuntu} alt="Ubuntu" />
                                          </div>
                                        </div>
                                      </div>
                                    </td>
                                )}
                                <td>{container.Name.split("-")[0]}</td>
                                <td>{container.Name.split("-")[1]}</td>
                            </tr>
                        ))    
                    )
                        
                )}
          </tbody>
        </table>
        </div>
        </>
    );
}