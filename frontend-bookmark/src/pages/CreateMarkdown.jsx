import React, { useState } from 'react';

import axios from 'axios';
import MDEditor from '@uiw/react-md-editor';
import Markdown from 'react-markdown';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomDark } from 'react-syntax-highlighter/dist/esm/styles/prism';

export default function CreateMarkdown() {
  const [value, setValue] = useState('**Hello world!!!**');
  const [savedValue, setSavedValue] = useState('');
  const [title, setTitle] = useState('');

  const handleEditorChange = (newValue) => {
    setValue(newValue);
    setSavedValue(newValue);
  };

  const handleTitleChange = (event) => {
    setTitle(event.target.value);
  };

  const saveToDatabase = () => {
    const createFile = {
      folder: 'new_folder',
      title: title,
      content: savedValue,
    };

    axios
      .post('http://localhost:8000/file/create', createFile)
      .then((response) => {
        console.log(response.data);
      })
      .catch((error) => {
        console.log(error);
      });

    console.log('Nilai yang disimpan:', savedValue);
    console.log('Nilai Title: ', title);
  };

  return (
    <div className="container">
      <h1>Welcome to Tauri!</h1>
      <input
        type="text"
        placeholder="Masukkan judul"
        value={title}
        onChange={(e) => handleTitleChange(e)}
      />
      <MDEditor value={value} onChange={handleEditorChange} />
      <Markdown
        children={value}
        components={{
          code(props) {
            const { children, className, node, ...rest } = props;
            const match = /language-(\w+)/.exec(className || '');
            return match ? (
              <SyntaxHighlighter
                {...rest}
                PreTag="div"
                children={String(children).replace(/\n$/, '')}
                language={match[1]}
                style={atomDark}
              />
            ) : (
              <code {...rest} className={className}>
                {children}
              </code>
            );
          },
        }}
      />
      <button onClick={saveToDatabase}>Simpan</button>{' '}
    </div>
  );
}
