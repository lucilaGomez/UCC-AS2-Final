import React from "react";
import "./HeroRegister.css";

const HeroRegister = ({ imageSrc }) => {
  return (
    <div className="hero">
      <img src={imageSrc} alt=" MyStay" className="hero__image" />
    </div>
  );
};

export default HeroRegister;