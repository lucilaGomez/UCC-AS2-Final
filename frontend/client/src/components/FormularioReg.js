import "./FormularioLogin.css";
import { useState } from "react";
import { FiMenu, FiX } from "react-icons/fi";
import { Link } from "react-router-dom";


export function FormularioReg ({setUser}){

    const [nombre, setNombre] = useState("")
    const [contrasena, setContrasena] = useState("")
    const [email, setEmail] = useState("")

    const [error, setError] = useState(false)

    const handleSubmit = (e) =>{

        e.preventDefault()
        if (nombre == "" || contrasena == "" || email == ""){
            setError(true)
            return
        }

        setError(false)

        setUser([nombre])

    }

    return(

        <section>

            <h1 className="Log">Ingrese sus datos:</h1>

            <form 
                className="formulario"
                onSubmit={handleSubmit}
            
            >

<input  
                placeholder="Nombre"
                type="text"
                value={nombre}
                onChange={e => setNombre(e.target.value)}
                ></input>

                <input  
                placeholder="Email"
                type="text"
                value={email}
                onChange={e => setEmail(e.target.value)}
                ></input>

                <input 
                placeholder="Contraseña"
                type="password"
                value={contrasena}
                onChange={e => setContrasena(e.target.value)}
                ></input>

                <button>Registrarse</button>


                
            </form>
            {error && <p className="camposOb">Todos los campos son obligatorios</p>}
        </section>
    )
}