export type SystemInfo = {
    CPU: {
        Used: string;
        Free: string;
    },
    Memory: {
        Used: string;
        Free: string;
    },
};

export type Container = {
    Name: string;
    ID: string;
}

export type Containers = {
    containers: Container[];
}