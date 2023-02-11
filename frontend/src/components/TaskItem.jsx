import React from "react";
import Button from "./UI/button/Button";

const TaskItem = (props) =>{
    if(props.task.active === props.setshow){
    return (
       
        <div className= 'task good'>
             <input
             type='checkbox'
             checked={!props.task.active}
             onChange = {event => 
                props.editTask(props.task,'active',!props.task.active)
            }
             />
        <div className='task_content'>
        <input
        type='text'
        value = {props.task.task}
        disabled = {!props.task.isred}
        onChange = {event => 
            props.editTask(props.task,'task',event.target.value)
        }
        />
        </div>
        <div className= 'task_buttons>'>
        <Button edit='delete' onClick = {() => props.removeTask(props.task)}>
          Удалить
        </Button>
        <Button edit='edit' onClick = {() => props.editTask(props.task,'isred',!props.task.isred)}>
          Изменить
        </Button>
        </div>
      </div>
    )
    }
}
export default TaskItem;