import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { atomDark } from 'react-syntax-highlighter/dist/esm/styles/prism';
import Markdown from 'react-markdown';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getPostById } from '@/service/post';

function PostsId() {
  const { id } = useParams();
  const [post, setPost] = useState({});

  useEffect(() => {
    getPostById(id)
      .then((post) => {
        console.log(post);
        setPost(post.data);
      })
      .catch((error) => console.log(error));
  }, []);

  const renderMarkdown = (markdown) => {
    console.log('Markdown:', markdown);
    if (markdown.Content) {
      const codeBlocks = markdown.Content.split('```');
      return (
        <div className="container mx-auto">
          <div className="w-full lg:w-8/12">
            <div className="mt-20 flex justify-center">
              <div className="bg-white p-6 rounded-lg shadow-lg w-full lg:w-8/12 mt-10">
                <h1 className="text-3xl font-bold">{markdown.Title}</h1>
                <p className="text-gray-500">
                  Date: 12 | Author: {markdown.User.Username}
                </p>
                {codeBlocks.map((block, index) => {
                  if (index % 2 === 1) {
                    return (
                      <SyntaxHighlighter
                        key={index}
                        language="javascript"
                        style={atomDark}
                        className="rounded-md p-4 mt-4 bg-gray-800 text-white"
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
                                className="rounded-md p-4 mt-4 bg-gray-800 text-white"
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
  };

  return (
    <div className="flex justify-center mt-10 h-screen">
      {renderMarkdown(post)}
    </div>
  );
}

export default PostsId;
