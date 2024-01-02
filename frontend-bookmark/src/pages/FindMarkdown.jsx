import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './css/findMarkdown.css';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomDark } from 'react-syntax-highlighter/dist/esm/styles/prism';
import Markdown from 'react-markdown';
import { useParams } from 'react-router-dom';

export default function FindMarkdown() {
  const { folder, filename } = useParams();
  const [markdowns, setMarkdowns] = useState({});

  const find = {
    folder: folder,
    fileName: filename + '.md',
  };

  useEffect(() => {
    axios
      .post('http://localhost:8000/file/find', find)
      .then((response) => {
        setMarkdowns(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  const renderMarkdown = (markdown) => {
    if (markdown.content) {
      const codeBlocks = markdown.content.split('```');
      return (
        <div className="container">
          <div className="row">
            <div className="col-md-8 offset-md-2">
              <div className="blog-post">
                <h1 className="blog-title">{markdown.title}</h1>
                <p className="blog-metadata">
                  Date: {markdown.date} | Author: {markdown.author}
                </p>
                {codeBlocks.map((block, index) => {
                  if (index % 2 === 1) {
                    return (
                      <SyntaxHighlighter
                        key={index}
                        language="javascript"
                        style={atomDark}
                      >
                        {block}
                      </SyntaxHighlighter>
                    );
                  } else {
                    return (
                      <Markdown
                        key={index}
                        children={block}
                        components={{
                          code(props) {
                            const { children, className, ...rest } = props;
                            const match = /language-(\w+)/.exec(
                              className || ''
                            );
                            return match ? (
                              <SyntaxHighlighter
                                {...rest}
                                language={match[1]}
                                style={atomDark}
                              >
                                {children}
                              </SyntaxHighlighter>
                            ) : (
                              <code {...props}>{props.children}</code>
                            );
                          },
                        }}
                      />
                    );
                  }
                })}
              </div>
            </div>
          </div>
        </div>
      );
    }
    return null;
  };

  return <div className="blog">{renderMarkdown(markdowns)}</div>;
}
