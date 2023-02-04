import React, {useState} from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

const MyCalendar = (props) =>{
    const [startDate, setStartDate] = useState(new Date(props.task.deadline));
    let dis = 'true'
    if(props.task.isred) dis=''
    return (
      <DatePicker
        selected={startDate}
        dateFormat="dd/MM/yyyy"
        onChange={
            (date) => {
              setStartDate(date)
              props.set(props.task, 'deadline', date)}
        }
        disabled = {dis}
      />
    );
}
export default MyCalendar