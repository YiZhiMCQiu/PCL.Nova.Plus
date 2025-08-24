# 0.0.3

1. 首次公开发布，用时两周
2. 目前更新自定义控件：MyDialog（自定义信息框）、MyLoading（自定义加载框）、MySelectCard（自定义下拉卡片）等一系列自定义组件
3. 更新了黑暗模式

# 0.0.4

1. 更新了MyToggleSwitch（自定义拨动开关）【应用至暗色模式的调整】
2. 更新了扫雷游戏！（更多 -> 小游戏 -> 扫雷），完美仿刻于：[xphost的Godot扫雷游戏](https://github.com/xphost008/Godot-Minesweeper)
3. 目前该扫雷缺陷是：仅仿刻于0.0.3版本，并且没有种子功能，同时游戏开局可能会直接踩雷直接输（因为开局就已经确定了所有雷的位置。。

# 0.0.5

1. 更换架构至`wails`，由于相较之前并没有什么新增内容，因此暂不发布二进制文件。。

# 0.0.6

1. 更新了MyPrograssBar控件！
2. 在主界面新增了【测试进度条】框，现在可以点击【开始】，然后进行测试加载了！

# 0.0.7

1. 在主界面的MyPrograssBar底下的文字新增了【进度显示】
2. 更新了MyPrograssBar的深色模式！
3. 更新了2048小游戏！可以去玩了！

# 0.0.8

1. MyToggleSwitch控件应用暗色模式！
2. 扫雷游戏新增了【扣分模式】，此时扫雷如果踩中雷不会扣分，而是减50分。
3. 2048游戏新增【作弊生成】按钮，按下即可生成256、512、1024三种随机数字在场上。ps：该操作不扣分也不加分，不减步数也不加步数。

# 0.0.9

1. 修复了扫雷无论点不点击扣分模式，都会扣分的bug
2. 将外置登录提上日程！！

# 0.0.10

1. 更了个新的控件：MyRadioButton！
2. 现在标题栏在从亮色切换到暗色时，亮度会稍微变暗一点。
3. 现在，暗色模式会保存到外部文件了！第一次启动时，会生成`{exe}\PCL.Nova\config\PCL2.Nova.ini`文件，可以保存一些配置信息。
4. 新增关于与鸣谢啦~【添加了原作者、现作者的鸣谢人员等】

# 0.0.11

1. 翻新账号部分！现在终于有CE那味了~~（虽然现在还暂时无法登录微软，不过可以登录离线登录啦！！）
2. 调整了一下信息框的背景颜色，现在看起来应该会更柔和了一点吧~
3. 稍微调整了一下加载框的暗色模式颜色。现在中间的镐子动画应该更加明显了。
4. 为NormalButton新增了disabled属性！现在可以正常显示禁用的样式啦~（不过正常的话需要手动判断此时是否处在禁用状态。。）
5. 账号部分目前可以正常保存到外部文件了！【别轻易发给别人（】
6. 为Nova添加了exe的图标啦！

# 0.0.12

1. 修改了MyLoading的样式，现在成功和失败时的左下角都会有一个图标了。
2. 修正了一下MySelectCard组件，现在已经不用maxHeight了~ （[PR#27](https://github.com/PCL-Community/PCL2.Nova.App/pull/27)，[@AMagicPear](https://github.com/AMagicPear)）
3. 新组件：MyCheckButton！（同时修改原来的MyCheckButton名称为MyNavButton）【与MyNavButton不同的是，该button选中时中间会打勾~】
   1. 👆感觉完全可以被MyToggleSwitch代替（但是又转念一想，好像也可以（目前MyCheckButton仅在主界面有展示，在别的地方目前没有实用空间。。
4. 初步设置透明度，并且现在已经支持设置背景图片啦！只需要将背景图片放到`{exe}\PCL.Nova\BackgroundImage`文件夹下，即可随机从中抽取一张背景图片展示！！
5. 现在开始，所有**组件**均有透明度127了！不用担心遮挡住背景图片了！
6. 稍微调亮了一点加载框成功时的颜色（
7. 在关于与鸣谢部分新增了PCL Community!
8. 将鸣谢改成特别鸣谢~

# 0.0.13 【以下由小万泥、xphost合作更新~】

## 将更新日志按照标准github语义来书写~

### refactor:

1. refactor(frontend): 将前端架构换成了svelte，以进一步减小打包体积。

### feat:

1. feat(Background): 翻新了一次背景样式，现在开始，如果未加入背景图片，则背景样式的渐变已经与原版PCL、CE一致啦~
2. feat(left): 重新制作了一次左侧栏的拉取，现在左侧栏已经和原版PCL类似了！切换页面时将不会收回再展开了！将会按照原版一样自动撑开~
3. feat(Thirdparty): 新增了第三方登录的初步界面，为以后添加第三方登录提供了一个接口~
4. feat(Minesweeper, P2048): 暂时移除了扫雷、2048等游戏，以后再慢慢加回来~
5. feat(TreasureBox): 新增百宝箱：内有今日人品和千万别点~

### style:

1. style(name): 名称已更换成PCL.Nova.Plus【不再使用原有的PCL2.Nova.App】

# 0.0.14

### fix:

1. fix(OfflineLogin): 修复了一个bug，无法自动通过账号名称绑定UUID到输入框的bug（未显示，但是能登录成功。。）
2. fix(OfflineLogin): 修复了名字可以命名超过3-16个字符的bug。。
3. fix(Card's Inner Margin): 稍微调整了一下所有有关于卡片内控件的外边距，现在已经变得更加贴合卡片了！
4. fix(About): 重新修改了帮助界面，现在帮助界面的按钮全部采用组件的形式了~但、但是缺点就是右边的Github主页变成icon形式的了。。
5. fix(MessageBox): 修复了当窗口移动到特别靠右的位置时，如果此时触发了信息框，很可能拉不回来。。修复方式为：现在信息框的背景是可以拖拽窗口的了！

### feat:

1. feat(Thirdparty): 在外置登录上，添加了【应用 Littleskin 服务器】的按钮，与此同时，如果 Littleskin 将来需要采用设备验证流方式登录，这边也会同样开始适配~
   1. 目前外置登录依旧不能用，因为此时MC都无法启动（所以登录的话即使做好了目前也暂时不会应用上去。。

### refactor:

1. refactor(Download): 初步重构下载部分，现在仅能获取 Minecraft 版本列表和 Forge、Fabric 的手动安装列表。但是无法下载，请见谅~
   1. 默认官方源，目前还没有做适配 BMCLAPI 的部分。下个版本把 Quilt、NeoForge、Optifine、LiteLoader 四个版本适配了之后，再做适配 BMCLAPI 的吧~

# 0.0.15

### refactor:

1. refactor(Launch): 已经可以启动游戏了！
   1. 目前，版本文件夹的选择是在【版本选择】里面，包括版本也是。点击之后的行为也与 PCL 一致。但是顶部菜单栏不会变（
   2. 版本独立设置目前暂时还没有，请见谅（
   3. 版本全局设置在【设置 -> 启动】里，这里也同样可以设置 Java
   4. Java 管理目前已被重构，重构得非常好看（确信）
2. refactor(Online): 现在可以在启动器中使用【IPv6】进行联机，
   1. 启动器支持以下 IPv6 功能：
      1. 检测本机 IPv6
      2. 尝试 ping 一次 IPv6
      3. 列举出本局域网 IPv6 是**临时**还是**永久**
      4. 输入游戏内端口，点击按钮复制 IPv6 地址与端口一同发送给朋友进行联机~
   2. 本启动器承诺绝对不会搜集您电脑本身的 IPv6 地址，如有疑问，请查阅源码。
   3. 本启动器检测 IPv6 的行为将与 LLL 启动器一致：[查看 LLL 源码](https://github.com/xphost008/lllauncher/blob/master/LittleLimboLauncher/OnlineIPv6Method.pas)

### feat:

1. feat(Controls): 更新了四个控件：
   1. MyInputBox: 用来显示输入框的
   2. MyNormalHint: 用来显示一个左下角提示。
   3. MyProgressBar: 滑动条组件
   4. MyCardButton: 在帮助页面、下载界面中显示的默认【类 PCL2 按钮】（
2. feat(GameSetting): 在 设置 -> 游戏 中新增了各种游戏设置，目前你已经可以正常的去玩了！
3. feat(about): 在 更多 -> 关于与鸣谢 中，明确指出了引用 MMCLL 的部分内容！

### style:

1. style(Left): 略微调整了主界面的左侧栏的宽度。
2. style(AccountSelect): 账号选择部分不再使用border来代替选中，而是在前方加上 MyRadioButton！

# 0.0.15.1

### fix:

1. fix(Java): 修复了在选择 Java 的时候，在没有选择任何 Java，第一次进入设置界面并生成 JavaJson.json 配置文件时，会弹出一个本不应该出现的报错（

### style:

1. style(messagebox): 略微调整了一下信息框的 padding

# 0.0.15.2

### fix:

1. fix(cmd): 稍微优化了一下执行终端的逻辑，现在开始，所有执行终端的输出，都将会静默执行，不会再显示终端的窗口了。（仅 Windows 的 cmd 的终端）
2. fix(Java): 修复了 Java 在前端展示信息时，误把路径写成版本号的 bug（
3. fix(IPv6): 稍微修复了一丢丢的 IPv6 获取逻辑（现在**应该**不会再导致窗口卡死了）

**# 0.0.16

### refactor:

1. refactor(HomePage): 制作了主页功能！各位可以在【设置 -> 主页】里面，输出一个自定义主页的教程！你也可以联网更新自己的自定义主页网址~
   - 自定义主页的教程与规范在[这里](./HomePageStandard.md)
   - 目前还没有支持 HTML 主页（呜呜呜，也许下个版本支持了呢？）
   - HTML 主页目前我还在纠结之中，是否应该支持，如果支持的话，会有什么安全隐患。。接下来就看用户反馈了吧~
2. refactor(Mod): 新增 Modrinth 的 Mod 搜索界面喵，现在已经可以正常搜索了喵！~

### exception:

1. 在 Config.go 文件里新增了 ExceptionHandler 的类型，用于在反馈给前端时提交一些重要信息~

### style:

1. style(MessageBox): 调整了信息框的 `rotate`，现在应该比 `PCL2` 更可爱了~
2. style(DontClick): 为【千万别点】新设计了两种，分别是单独将 `rotateX` 和 `rotateY` 转换成 `180deg`，仅此而已（
3. style(MyNormalSpan): 将 `MyNormalLabel` 修改成了 `MyNormalSpan`，并且将内部元素也改成了 `span` 标签，现在应该更加合理了！
4. style(Controls): 将几乎所有的控件的点击事件，全部设计成了 dispatch！现在应该可以在使用该控件位置的地方使用【on:事件名】来调用了！

### fix:

1. fix(afdian): 修复了爱发电图标失效的bug
2. fix(IPv6): 彻底修复了 IPv6 检测~
3. fix(MCVersion): MC 文件夹管理可能会导致一些问题
4. fix(AccountSelect): 启动游戏时可能出**现的 Cannot convert account index to int 报错~