import React from "react";
import { Link } from "react-router-dom";
    
    
    const navbarLinks = [

    { url: "/", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },
      ];
      
      
      
      const DisponibilidadnReserva = () => {
        return (
          <div className="confirmacion-reserva">
                <Link to="/">
              <button>Volver al menu principal</button>
            </Link>
          </div>
          
        );
      };
      
      export default DisponibilidadnReserva;
      
