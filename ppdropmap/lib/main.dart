import 'package:flutter/material.dart';
import 'package:google_maps/google_maps.dart' hide Icon;
import 'package:url_launcher/url_launcher.dart';
// ignore: avoid_web_libraries_in_flutter
import 'dart:html';
import 'dart:ui' as ui;

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'pp-drop map',
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
    if (await canLaunch(_link)) {
      await launch(_link);
    } else {
      throw 'Could not launch $_link';
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: <Widget>[
          _map(),
          _redirect(),
        ],
      ),
    );
  }

  Widget _map() {
    String htmlId = "7";

    // ignore: undefined_prefixed_name
    ui.platformViewRegistry.registerViewFactory(htmlId, (int viewId) {
      final mc1 = LatLng(50.4510788, 30.5192703);
      final mc2 = LatLng(50.4491999, 30.5226107);

      final mapOptions = MapOptions()
        ..zoom = 16
        ..clickableIcons = true
        ..disableDefaultUI = true
        // ..streetViewControl = false
        // ..zoomControl = false
        ..center = mc2;

      final elem = DivElement()
        ..id = htmlId
        ..style.width = "100%"
        ..style.height = "100%"
        ..style.border = 'none';

      final map = GMap(elem, mapOptions);

      Marker(MarkerOptions()
        ..position = mc1
        ..map = map
        ..clickable = true
        ..title = 'Hello');

      Marker(MarkerOptions()
        ..position = mc2
        ..map = map
        ..clickable = true
        ..title = 'Kyiv');

      return elem;
    });

    return HtmlElementView(viewType: htmlId);
  }

  Widget _redirect() {
    final color = Color(0xff2ecc71);
    return Padding(
      padding: EdgeInsets.all(24),
      child: Align(
        alignment: Alignment.bottomRight,
        child: FlatButton(
          visualDensity: VisualDensity.comfortable,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(24),
          ),
          padding: EdgeInsets.symmetric(horizontal: 28, vertical: 16),
          clipBehavior: Clip.antiAlias,
          color: color,
          focusColor: color,
          highlightColor: Colors.black,
          hoverColor: Colors.black.withOpacity(1),
          child: Row(
            mainAxisSize: MainAxisSize.min,
            children: <Widget>[
              Text(
                'pp-drop',
                style: TextStyle(fontSize: 24, color: Color(0xffecf0f1)),
              ),
              Icon(
                Icons.arrow_forward_ios,
                color: Color(0xffecf0f1),
              ),
            ],
          ),
          onPressed: _onPressed,
        ),
      ),
    );
  }
}
