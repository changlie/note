PCM数据：它是模拟音频信号经模数转换（A/D变换）直接形成的二进制序列，该文件没有附加的文件头和文件结束标志。  

声音本身是模拟信号，而计算机只能识别数字信号，要在计算机中处理声音，就需要将声音数字化，这个过程叫经模数转换（A/D变换）。最常见的方式是透过脉冲编码调制PCM(Pulse Code Modulation) 。

描述声音的四个维度：
- **采样频率：** 即取样频率,指每秒钟取得声音样本的次数。采样频率越高,声音的质量也就越好,声音的还原也就越真实，但同时它占的资源比较多。由于人耳的分辨率很有限,太高的频率并不能分辨出来。在16位声卡中有22KHz、44KHz等几级,其中，22KHz相当于普通FM广播的音质，44KHz已相当于CD音质了，目前的常用采样频率都不超过48KHz。
- **采样位数：** 即采样值或取样值（就是将采样样本幅度量化）。它是用来衡量声音波动变化的一个参数，也可以说是声卡的分辨率。它的数值越大，分辨率也就越高，所发出声音的能力越强。
- **声道数：** 很好理解，有单声道和立体声之分，单声道的声音只能使用一个喇叭发声（有的也处理成两个喇叭输出同一个声道的声音），立体声的PCM 可以使两个喇叭都发声（一般左右声道有分工） ，更能感受到空间效果。
- **时长：** 采样的时长
> 采样频率单位为Hz，表示每秒采样的次数，一般有11025HZ（11KHz），22050HZ（22KHz）、44100Hz（44KHz）三种。  
> 采样位数单位为bit（位），一般有8bit和16bit 。8bit表示用8bit空间量化某时刻的声音，这一点基本是和图片用r、g、b三单位共24bit量化颜色一样。
- `44KHz,16bit`的声音称作：CD音质
- `22KHz,16bit`的声音效果近似于立体声（FM Stereo）广播, 称作：广播音质
- `11kHz,8bit`的声音, 称作：电话音质
