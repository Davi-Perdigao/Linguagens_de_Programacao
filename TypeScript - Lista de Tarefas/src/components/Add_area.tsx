import {useState, KeyboardEvent} from 'react'
import styled from 'styled-components'

const Container_1a = styled.div`
border: 1px solid #555;
border-radius: 15px;
padding: 10px;
margin: 20px 0;
display: flex;
align-item: center;`

const Div_1a = styled.div`
margin-right 10px;`

const Input_1a = styled.input`
border: 0;
background: transparent;
outline: 0;
color: #fff;
font-size: 18px;
flex: 1`

type Props = {
  onEnter: (taskName: string) => void
}

export default ({onEnter}: Props) => {
  const [inputText, setInputText] = useState('')

  const handleKeyUp = (e: KeyboardEvent) => {if(e.code === 'Enter' && inputText !== '') {
    onEnter(inputText);
    setInputText('');
  }}

  return(
    <Container_1a>
      <Div_1a>➕
      </Div_1a>
      <Input_1a
      type='text'
      placeholder='Adcione uma terefa'
      value={inputText}
      onChange={e => setInputText(e.target.value)}
      onKeyUp={handleKeyUp}
      />
    </Container_1a>
  )
}