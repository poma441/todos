import React from "react"

const defaultState = {
    tasks: [
      {id: 1, task:'make front',active: true,  isred: false},
      {id:2, task: 'make back',active: true,  isred: false},
      {id:3, task: 'make back2',active: false,  isred: false},
      {id:4, task: 'make back3',active: false,  isred: false}
    ]
  }


export const taskReducer = (state = defaultState, action) => {
    switch (action.type){
      case "ADD_TASK":
            console.log(action.payload)
            return {...state, tasks:[...state.tasks, action.payload]}
      case "REMOVE_TASK":
            return {...state, tasks: state.tasks.filter(tsk=>tsk.id!==action.payload.id)}
      
      case "EDIT_TASK":
            return  {...state, tasks: state.tasks.map(tsk => {
               if (tsk.id === action.payload.task.id) return {...tsk, [action.payload.param]:action.payload.value }
               else return tsk;
              })
            }

      default:
        return state
    }

  }