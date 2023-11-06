import React from "react"
import { useNavigate } from "react-router"
import Navbar from "./components/Navbar";
import "./DatePicker.css";


export default function MisReservas(){


    const navigate = useNavigate();
    const navbarLinks = [

    { url: "/", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },
      ];

    return (<div className="misReservas">

        <Navbar navbarLinks={navbarLinks} />

            <div className="date-container">

               
                
            </div>

    </div>
    )

}