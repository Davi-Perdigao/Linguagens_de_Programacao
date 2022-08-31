import React from 'react';
import {useState} from 'react';
import styled from 'styled-components';
import {Item} from './types/Item';
import List_item from './components/List_item'
import Add_area from './components/Add_area';

const Container_1a = styled.div`
background-color: #17181F;
color: #797A81;
min-height: 100vh;`

const Container_2a = styled.div`
margin: auto;
max-width: 980px;
padding: 10px;`

const Title_1a = styled.h1`
margin: 0;
padding: 0;
color: #ffffff;
text-align: center;border-bottom: 1px solid #444;
padding-bottom: 20px;
`

export default () => {
  const [list, setList] = useState<Item[]>([]);

  const handleAddTask = (taskName: string) => {
    let newList = [...list];
    newList.push({
      id: list.length + 1,
      name: taskName,
      done: false
    });
    setList(newList);
  }

  const handleCheck = (taskId: number) => {
    const result = list.findIndex(i => i.id === taskId);
    if (list[result].done === false) {list[result].done = true}
    else {list[result].done = false};
  }

  return (
    <Container_1a>
      <Container_2a>
        <Title_1a>
          Lista de Tarefas
        </Title_1a>
        <Add_area onEnter={handleAddTask}/>
        {list.map((item, index)=>(
          <List_item key={index} item={item} onCheck={handleCheck}/>
        ))}
      </Container_2a>
    </Container_1a>
  )
}