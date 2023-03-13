import { Box, Flex } from "@chakra-ui/react";
import axios from "axios";
import { useRef, useState, useCallback, SetStateAction, Dispatch, memo, VFC } from "react";
import Webcam from "react-webcam";
import "../styles.css";

type Props = {
  setContent: Dispatch<SetStateAction<string>>;
  content: string;
}

// 参考：https://qiita.com/ko-izumi/items/060c258b16a9ed1294c6

export const WebCam: VFC<Props> = memo((props) => {
  const {setContent} = props;
  const [isCaptureEnable, setCaptureEnable] = useState<boolean>(false);
  // WebCamパッケージのWebCamオブジェクトをuseRefでインスタンス化し、webcamRefに格納
  const webcamRef = useRef<Webcam>(null);
  const [url, setUrl] = useState<string | null>(null);
  const capture = useCallback(() => {
    const imageSrc = webcamRef.current?.getScreenshot();
    if (imageSrc) {
      setUrl(imageSrc);
    }
  }, [webcamRef]);

  // 画像の送信処理 Todo:関数化(hook化)したい
  const handleUpload = async () => {
    // const imageSrc = webcamRef.current?.getScreenshot();
    // if (!imageSrc) {
    //   console.error('Failed to capture image');
    //   return;
    // }
    // const formData = new FormData();
    // formData.append("image", imageSrc);
    // console.log(formData)

    const imageSrc = webcamRef.current?.getScreenshot();
    if (!imageSrc) {
      console.error('Failed to capture image');
      return;
    }
    const byteCharacters = atob(imageSrc.split(',')[1]); // base64 データの取り出し
    const byteNumbers = new Array(byteCharacters.length);
    for (let i = 0; i < byteCharacters.length; i++) {
      byteNumbers[i] = byteCharacters.charCodeAt(i);
    }
    const byteArray = new Uint8Array(byteNumbers);
    const blob = new Blob([byteArray], { type: 'image/jpeg' });

    const formData = new FormData();
    formData.append("image", blob, 'image.jpg');
  
    try {
      const response = await axios.post('http://localhost:8080/convertImage', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
        responseType: 'blob'
      });

      const reader = new FileReader();
      reader.onload = () => {
        const decoder = new TextDecoder('utf-8');
        const str = decoder.decode(new Uint8Array(reader.result as ArrayBuffer));
        console.log(str)
        setContent(str);
      };
      reader.readAsArrayBuffer(response.data);
    } catch (err) {
      console.error(err);
    }
  };
  

  // フロントカメラとの切り替え
  const [facingModeState, setFacingMode] = useState('user');
  const handleFacingModeToggle = () => {
    setFacingMode((prevMode) => (prevMode === 'user' ? 'environment' : 'user'));
  };

  const videoConstraints = {
    width: 720,
    height: 360,
    facingMode: facingModeState,
  };

  return (
    <>
      <header>
        <h1>カメラアプリ</h1>
      </header>
      {isCaptureEnable || (
        <button onClick={() => setCaptureEnable(true)}>開始</button>
      )}
      {isCaptureEnable && (
        <>
          <Flex align="center" justify="center" height="10vh">
            <Box>
              <div>
                <Webcam
                  audio={false}
                  width={540}
                  height={360}
                  ref={webcamRef}
                  screenshotFormat="image/jpeg"
                  videoConstraints={videoConstraints}
                />
              </div>
          
          <button onClick={capture}>キャプチャ</button>
          <br/>
          <button onClick={handleFacingModeToggle}>
            {facingModeState === 'user' ? 'リアカメラに切り替え' : 'フロントカメラに切り替え'}
          </button>
          <div>
            <button onClick={() => setCaptureEnable(false)}>終了</button>
          </div>
          </Box>
          <br/>

          
          </Flex>
          
        </>
      )}
      {url && (
        <>
          <div>
            <button
              font-color="red"
              onClick={() => {
                setUrl(null);
              }}
            >
              削除
            </button>
          </div>
          <br/>
        
          <Box>
            <img src={url} />
            <br/>
            <button font-color="red" name="image" id="image" onClick={handleUpload}>Upload photo</button>
            <br/>
            <input type="file" onChange={handleUpload} />
          </Box>
        
        </>
      )}
    </>
  );
});