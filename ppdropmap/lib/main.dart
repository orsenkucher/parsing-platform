import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:url_launcher/url_launcher.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'pp-drop: map'),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatefulWidget {
  final String title;
  MyHomePage({Key key, this.title}) : super(key: key);

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  String _link = "https://t.me/ppdropbot";

  void _onPressed() async {
    // setState(() {});
    // if (await canLaunch(_link)) {
    //   await launch(_link);
    // } else {
    //   throw 'Could not launch $_link';
    // }
    // SystemChannels.platform.invokeMethod('SystemNavigator.pop');
    try {
      SystemNavigator.pop();
    } on dynamic catch (err) {
      print(err);
    }
  }

  @override
  Widget build(BuildContext context) {
    final color = Colors.blue;
    return Scaffold(
      body: Center(
        child: OutlineButton(
          borderSide: BorderSide(
            color: color,
            width: 8,
            style: BorderStyle.solid,
          ),
          visualDensity: VisualDensity.comfortable,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(24),
          ),
          padding: EdgeInsets.all(24),
          clipBehavior: Clip.antiAlias,
          color: color,
          focusColor: color,
          hoverColor: color.withOpacity(0.2),
          child: Text(
            'Open pp-drop',
            style: TextStyle(fontSize: 36),
          ),
          onPressed: _onPressed,
        ),
      ),
    );
  }
}
