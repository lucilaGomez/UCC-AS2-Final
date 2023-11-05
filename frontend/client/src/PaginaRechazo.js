import React from "react";
import { Link } from "react-router-dom";
    
    
    const navbarLinks = [

    { url: "/", title: "Reservar" },
    { url: "/misreservas", title: "Mis Reservas" },
    { url: "/login", title: "Login" },
      ];
      
      
      
      const RechazoReserva = () => {
        return (
          <div className="Rechazo-reserva">
            <h2>Lo siento, su reserva ha sido denegada</h2>
            <Link to="/">
              <button>Volver al menu principal</button>
            </Link>
          </div>
        );
      };
      
      export default RechazoReserva;
      
