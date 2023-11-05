import "./FormularioLogin.css";
import { useState } from "react";
import { useNavigate } from "react-router"


export function FormularioLogin ({setUser}){

    const [nombre, setNombre] = useState("")
    const [contrasena, setContrasena] = useState("")
    const [error, setError] = useState(false)


    const handleSubmit = (e) =>{

        e.preventDefault()
        if (nombre == "" || contrasena == ""){
            setError(true)
            return
        }

        setError(false)

        setUser([nombre])

    }


    const navigate = useNavigate();

    return(

        <section>

            <h1 className="Log">Login</h1>

            <form 
                className="formulario"
                onSubmit={handleSubmit}
            
            >

                <input  
                placeholder="Usuario"
                type="text"
                value={nombre}
                onChange={e => setNombre(e.target.value)}
                ></input>

                <input 
                placeholder="Contraseña"
                type="password"
                value={contrasena}
                onChange={e => setContrasena(e.target.value)}
                ></input>

                <button>Iniciar Sesión</button>

                <button onClick={()=>navigate("/register")}>No tiene cuenta? Registrese!</button>
               
                

            </form>
            {error && <p className="camposOb">Todos los campos son obligatorios</p>}
        </section>
    )
}