import { materialCells, materialRenderers } from '@jsonforms/material-renderers';
import { JsonForms } from '@jsonforms/react';
import { useEffect, useState } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { monokai } from 'react-syntax-highlighter/dist/esm/styles/hljs';

function USCDLEditor() {
  const [schema, setSchema] = useState(null);
  const [data, setData] = useState(null);
  const [generatedFiles, setGeneratedFiles] = useState([]);
  const [selectedFile, setSelectedFile] = useState(null);
  const [fileContent, setFileContent] = useState('');
  const [uiSchema, setUiSchema] = useState({
    type: 'VerticalLayout',
    elements: [
      {
        type: 'Control',
        scope: '#/properties/containers'
      }
    ]
  });

  useEffect(() => {
    // Load the schema
    fetch('http://localhost:8080/schema.json')
      .then(response => response.json())
      .then(schemaData => {
        setSchema(schemaData);
      })
      .catch(error => console.error('Error loading schema:', error));

    // Load sample data
    fetch('http://localhost:8080/adcs.json')
      .then(response => response.json())
      .then(jsonData => {
        setData(jsonData);
      })
      .catch(error => console.error('Error loading data:', error));

    // Load generated files
    fetchGeneratedFiles();
  }, []);

  const fetchGeneratedFiles = () => {
    fetch('http://localhost:8080/generated')
      .then(response => response.json())
      .then(files => {
        setGeneratedFiles(files);
        if (files.length > 0 && !selectedFile) {
          setSelectedFile(files[0]);
          fetchFileContent(files[0].name);
        }
      })
      .catch(error => console.error('Error loading generated files:', error));
  };

  const fetchFileContent = (filename) => {
    fetch(`http://localhost:8080/generated/${filename}`)
      .then(response => response.text())
      .then(content => {
        setFileContent(content);
      })
      .catch(error => console.error(`Error loading file ${filename}:`, error));
  };

  const handleFormChange = ({ data: newData }) => {
    setData(newData);
  };

  const downloadJson = () => {
    const jsonString = JSON.stringify(data, null, 2);
    const blob = new Blob([jsonString], { type: 'application/json' });
    const url = URL.createObjectURL(blob);

    const a = document.createElement('a');
    a.href = url;
    a.download = 'spacecraft_data.json';
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  const generateCode = () => {
    fetch('http://localhost:8080/generate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then(response => response.json())
      .then(result => {
        alert('Code generation successful!');
        fetchGeneratedFiles(); // Refresh the file list
      })
      .catch(error => {
        console.error('Error generating code:', error);
        alert('Failed to generate code. See console for details.');
      });
  };

  const getLanguage = (file) => {
    if (!file) return 'text';
    if (file.type === 'c') return 'c';
    if (file.type === 'h') return 'cpp'; // C headers can use C++ highlighting
    if (file.type === 'ts') return 'typescript';
    return 'text';
  };

  if (!schema || !data) {
    return <div className="container mt-5"><div className="alert alert-info">Loading schema and data...</div></div>;
  }

  return (
    <div className="container-fluid">
      <div className="row mb-4">
        <div className="col-12">
          <h1 className="my-4">Spacecraft Data Definition Editor</h1>
        </div>
      </div>

      <div className="row mb-4">
        <div className="col-12">
          <div className="card">
            <div className="card-body">
              <h5 className="card-title mb-3">Data Structure Editor</h5>
              <JsonForms
                schema={schema}
                uischema={uiSchema}
                data={data}
                renderers={materialRenderers}
                cells={materialCells}
                onChange={handleFormChange}
              />
            </div>
          </div>
        </div>
      </div>

      <div className="row mb-4">
        <div className="col-12">
          <div className="btn-group">
            <button onClick={downloadJson} className="btn btn-primary">
              Download JSON
            </button>
            <button onClick={generateCode} className="btn btn-success ms-2">
              Generate Code
            </button>
          </div>
        </div>
      </div>

      <div className="row mb-4">
        <div className="col-12">
          <div className="card">
            <div className="card-body">
              <h5 className="card-title">JSON Preview</h5>
              <SyntaxHighlighter language="json" style={monokai} className="rounded">
                {JSON.stringify(data, null, 2)}
              </SyntaxHighlighter>
            </div>
          </div>
        </div>
      </div>

      {generatedFiles.length > 0 && (
        <div className="row mb-4">
          <div className="col-12">
            <div className="card">
              <div className="card-body">
                <h5 className="card-title mb-3">Generated Files</h5>
                <div className="row">
                  <div className="col-md-3">
                    <div className="list-group">
                      {generatedFiles.map(file => (
                        <button
                          key={file.name}
                          className={`list-group-item list-group-item-action ${selectedFile && selectedFile.name === file.name ? 'active' : ''}`}
                          onClick={() => {
                            setSelectedFile(file);
                            fetchFileContent(file.name);
                          }}
                        >
                          {file.name}
                        </button>
                      ))}
                    </div>
                  </div>
                  <div className="col-md-9">
                    {selectedFile && (
                      <SyntaxHighlighter
                        language={getLanguage(selectedFile)}
                        style={monokai}
                        className="rounded"
                        showLineNumbers={true}
                      >
                        {fileContent}
                      </SyntaxHighlighter>
                    )}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default USCDLEditor;