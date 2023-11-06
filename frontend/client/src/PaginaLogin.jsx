import Navbar from "./components/Navbar";
import loginphoto from "./assets/loginphoto.jpg";
import { FormularioLogin } from "./components/FormularioLogin";
import {Home} from './components/Home';
import { useState } from "react";
import HeroLogin from "./components/HeroLogin";


export default function PaginaLogin(){

    const [user, setUser] = useState([])

    const navbarLinks = [

    { url: "/", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },
      ];

      

    return <div className="PaginaLogin">
        
        <Navbar navbarLinks={navbarLinks} />
        <HeroLogin imageSrc={loginphoto} />

        {
            !user.length > 0
            ?<FormularioLogin setUser={setUser}/>
            :<Home user={user}/>
        }

        
        </div>
}