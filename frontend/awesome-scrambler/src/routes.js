import Home from "./components/Home";
import CipherText from "./components/CipherText";


export const routes = [
    {
        path: "/",
        component: Home,
    },
    {
        path: "/ciphertext/:path",
        component: CipherText,
    },
]