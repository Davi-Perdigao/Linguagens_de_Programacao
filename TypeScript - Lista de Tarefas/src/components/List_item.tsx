import {useState} from 'react';
import styled from 'styled-components';
import {Item} from '../types/Item'

type Props = {
  item: Item;
  onCheck: (taskId: number) => void
}

type Label_1aProps = {
  done: boolean;
}

const Container_1a = styled.div`
display: flex;
background-color: #20212C;
padding: 10px;
border-radius: 10px;
margin-bottom: 10px;
align-item: center;
position: relative`

const Input_1a = styled.input`
width: 20px;
height: 20px;
margin-right: 10px;
cursor: pointer`

const Label_1a = styled.div(({done}: Label_1aProps)=>(
`
color: #ccc;
text-decoration: ${done ? 'line-through' : 'initial'}
`
));

export default ({item, onCheck}: Props) => {
  const [isChecked, setIsChecked] = useState(item.done);

  return (
    <Container_1a>
      <Input_1a
      type="checkbox"
      checked={isChecked}
      onChange={e => setIsChecked(e.target.checked)}
      onClick={() => onCheck(item.id)}/>
      <Label_1a done={isChecked}>{item.name}</Label_1a>
    </Container_1a>
  );
}