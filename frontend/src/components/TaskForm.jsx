import React, {useState} from "react";
import CustomInput from "./UI/input/CustomInput";
//import Button from "./UI/button/Button";
import Button  from '@mui/material/Button' ;

const TaskForm = ({create}) => {

    const [task, setTask] = useState({id:'', task:'', isred: false, active: false});

    
    const addNewTask = (e) =>{
        e.preventDefault()
        const newTask = {
            ...task, 
            id: Date.now(),
            isred: false,
            active: true
        }
        create(newTask)
        setTask({task:'', isred: false})
    }

    return (
        <div>
             <form>
                <CustomInput
                value = {task.task}
                type='text' 
                placeholder='Задача'
                onChange={e=>setTask({...task, task:e.target.value})}
                />
                <Button onClick={addNewTask}>Создать задачу</Button>
            </form>
        </div>
    )
};
export default TaskForm;