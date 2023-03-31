import * as React from 'react';
import imgUrl from '../north_america_maps.jpg'
import sunsetwxUrl from '../sunrise_et.png'

function changeImageColor (imgUrl, color = "#ffbf00") {

    let threshold = 40; //默认颜色阀值 为 114 －>和默认图相关
    let img = new Image();
    img.src = imgUrl;
    let newR = parseInt("0x" + color.substr(1, 2));
    let newG = parseInt("0x" + color.substr(3, 2));
    let newB = parseInt("0x" + color.substr(5, 2));
    //图片加载后进行处理
    img.onload = function () {
        console.log("wojinlaile");
        let width = img.width, height = img.height,canvas = document.getElementById("mycanvas")
        // @ts-ignore
        let ctx = canvas.getContext("2d");
        // @ts-ignore
        canvas.width = width;
        // @ts-ignore
        canvas.height = height;
        // 将源图片复制到画布上
        if (ctx){


            ctx.drawImage(img, 0, 0, width, height);
            ctx.globalCompositeOperation = 'source-over';
            ctx.fillStyle = 'red';
            // 获取画布的像素信息
            let imageData = ctx.getImageData(0, 0, width, height),data = imageData.data;
            // 对像素集合中的单个像素进行循环，每个像素是由4个通道组成，所以要注意
            let i = 0;
            while (i < data.length/2) {
                let r = data[i++],
                    g = data[i++],
                    b = data[i++],
                    a = data[i++];
                //计算透明度
                let alp = (255 - r) / (255 - threshold);
                //判断是否透明
                let isTransparent = r == 255 && g == 255 && b == 255 && a == 255;
                if (isTransparent) {
                    data[i - 1] = 0;
                } else {
                    data[i - 4] = newR;
                    data[i - 3] = newG;
                    data[i - 2] = newB;
                    data[i - 1] = a !== 255 ? 255 - a : alp * 255; //处理透明的图片和不透明的图片
                }
            }
            // 将修改后的代码复制回画布中
            ctx.putImageData(imageData, 0, 0);
            // 图片导出为 png 格式
            let imgType = "png";
            // @ts-ignore
            let imgData = canvas.toDataURL(imgType);
            // alert(imgData); // 生成base64
            // console.log(imgData);

        }};
    return img;
}
//使用方式 其中 data、style 为通过api读取的值

export default function Mapshow() {

    // let Img = changeImageColor(imgUrl,"#ff4d4d")
    return (
        <div>
            {/*<canvas id='mycanvas' style={{*/}
            {/*    display: "block",*/}
            {/*    marginLeft: "auto",*/}
            {/*    marginRight: "auto",*/}
            {/*    textAlign: "center"*/}
            {/*}}>*/}

            {/*</canvas>*/}
        <img
            src="../sunrise_et.png"
            alt="My Image"
            style={{
            display: "block",
            marginLeft: "auto",
            marginRight: "auto",
            textAlign: "center"
        }}
            />
        </div>


    );
}
