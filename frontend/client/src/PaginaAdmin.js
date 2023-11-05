import React from "react"
import { useNavigate } from "react-router"

export default function PaginaAdmin(){


    const navigate = useNavigate();

    return (<div className="paginaAdmin">


            <form>
                <input type="text" placeholder="usuario admin"></input>
                <input type="password" placeholder="contraseña"></input>
                <button onClick={()=> navigate("/comandos")}>Iniciar Sesión</button>
            </form>

    </div>
    )

}