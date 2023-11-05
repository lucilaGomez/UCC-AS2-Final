import React, { useState } from "react";
import { DateRange } from "react-date-range";
import "react-date-range/dist/styles.css";
import "react-date-range/dist/theme/default.css";


const DatePicker = () => {
  //date state
  const [range, setRange] = useState([
    {
      startDate: new Date(),
      endDate: new Date(),
      key: "selection",
    },
  ]);

  const handleSelectDates = (ranges) => {
    setRange([ranges.selection]);
  };

  //determina la fecha de hoy
  const today = new Date();

  return (
    <div>
      <DateRange
        ranges={range}
        onChange={handleSelectDates}
        editableDateInputs={true}
        moveRangeOnFirstSelection={false}
        months={2}
        minDate={today}
        direction="horizontal"
        className="calendarElement"
     />
    </div>
    
  );
};


export default DatePicker;