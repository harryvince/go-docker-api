import { toast } from "react-toastify";
import { CreateContainer } from "../utils/api";

type props = {
    image: string,
    container_type: "mysql" | "ubuntu" | "postgres"
}

export default function Button(props: props) {
    const { image, container_type } = props;
    const createContainer = CreateContainer(container_type);

    // if (createContainer.data) {
    //     toast.info('🦄 Wow so easy!', {
    //         position: "bottom-right",
    //         autoClose: false,
    //         hideProgressBar: false,
    //         closeOnClick: false,
    //         pauseOnHover: true,
    //         draggable: false,
    //         progress: undefined,
    //         theme: "colored",
    //     });
    // }

    return (
        <>
            <button className="btn btn-lg m-2" onClick={() => createContainer.mutate()}>
                <img src={image} alt={image} height={25} width={25} />
            </button>
        </>
    );
}