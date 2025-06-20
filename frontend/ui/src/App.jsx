import React, { useState, useEffect,useRef } from 'react';
import './App.css';

export default function App() {
  const [expression, setExpression] = useState('');
  const [isValid, setIsValid] = useState(null);
  const editorReference = useRef();

  useEffect(() => {
    if (expression === '') {
      setIsValid(null);
      return;
    }

    fetch('http://localhost:8080/validate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ expression: expression })
    })
    .then(res => res.json())
    .then(data => {
      setIsValid(data.valid);
    })
    .catch(() => {
      setIsValid(false);
    });

  
  }, [expression]);

  const color = () => {
    if(isValid === null){
      return '#ffcc70'
    }
    else if(isValid===true){
      return '#66BB6A';
    }
    else{
      return '#EF5350';
    }
  }

  const colorChars = () => {
  
    const input = editorReference.current.textContent || '';
    
    //Deletes special chars that the browser usually leaves in an input, with this if you erase everything in the input, it can give a null expression and turn into yellow. The n doesn't allow you to jump line
    const expressionInput = input.replace(/[\u200B-\u200D\uFEFF\n]/g, '')
     
    setExpression(expressionInput);

    const colorSpan= expressionInput.split('').map((char) => {
  
      if ("+-*/".includes(char)) {
        return `<span class="operator">${char}</span>`;
      }
      else if("()".includes(char)){
        return `<span class="parenthesis">${char}</span>`;
      }
      else if("[]".includes(char)){
        return `<span class="corchetes">${char}</span>`;
      } 
      else if("0123456789".includes(char)){
        return `<span class="numbers">${char}</span>`;
      } 
      else {
        return `<span class="others">${char}</span>`;
      }

    }).join('');

    //this puts the color of the expression
    editorReference.current.innerHTML = colorSpan;

    //this moves the cursor of the input
    const range = document.createRange();
    const selected = window.getSelection();
    range.selectNodeContents(editorReference.current);
    range.collapse(false);
    selected.removeAllRanges();
    selected.addRange(range);
  }

  return (
    <div className="container" style={{ borderColor: color() }}>
    
      <h1 style={{ color: color() }}>Math Expression Analyzer</h1>
      <h2>Write your expression: </h2>
      
      <div className="inputContainer" style={{ borderColor: color() }}> 
        <div
            className="editor"
            ref={editorReference}
            contentEditable
            onInput={colorChars}
        ></div>
      </div>

      <div className="status">
        {isValid === null && <p className="empty"></p>}
        {isValid === true && <p className="valid"></p>}
        {isValid === false && <p className="invalid"></p>}
      </div>
    </div>
  )
}
