import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

const Inforeserva = () => {
  const [hotels, setHotels] = useState([]);

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const response = await axios.get("http://localhost:8080/hotel");
        setHotels(response.data);
      } catch (error) {
        console.error("Error al obtener datos del servidor:", error);
      }
    };

    fetchHotels();

  }, []);

  return (
    <div className="confirmacion-reserva">
      <div className="hotels-container">
        {hotels.length ? (
          hotels.map((hotel) => (
            <div className="hotel-card" key={hotel.id}>
              <h4>{hotel.nombre}</h4>
              <p>{hotel.email}</p>
              <p>Descripción: {hotel.descripcion}</p>
            </div>
          ))
        ) : (
          <p>No hay hoteles disponibles.</p>
        )} 
      </div>
      <Link to="/">
        <button>Volver al menú principal</button>
      </Link>
    </div>
  );
};

export default Inforeserva;
