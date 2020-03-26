// import 'package:easy_google_maps/easy_google_maps.dart';
// ignore: avoid_web_libraries_in_flutter
import 'dart:js';

import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

import 'package:google_maps/google_maps.dart' hide Icon;
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
      initialRoute: '/',
      // My url pattern: www.app.com/#/xLZppqzSiSxaFu4PB7Ui
      onGenerateRoute: (settings) {
        print("=========");
        // print(settings);
        // print('route: ' + settings.name);
        // print(settings.arguments);
        final uri = Uri.tryParse(settings.name);
        print('uri: $uri');

        // var params = <String, String>{};
        // try {
        //   params = uri?.queryParameters;
        // } catch (_) {
        //   print(_);
        // }
        final params = uri.queryParameters;
        // print(params);
        params.forEach((k, v) {
          print('-> $k: $v');
        });
        // print(params['chatid']);
        // print(params['lat']);
        // print(params['lng']);
        // print("=========");

// TODO
        return MaterialPageRoute(builder: (context) {
          return MyHomePage(
            chatid: params['chatid'],
            lat: params['lat'] != null ? double.tryParse(params['lat']) : null,
            lng: params['lng'] != null ? double.tryParse(params['lng']) : null,
          );
        });

        // List<String> pathComponents = settings.name.split('/');
        // if (pathComponents[1] == 'invoice') {
        //   return MaterialPageRoute(
        //     builder: (context) {
        //       // return Invoice(arguments: pathComponents.last);
        //       return Placeholder();
        //     },
        //   );
        // } else {
        //   return MaterialPageRoute(
        //     builder: (context) {
        //       // return LandingPage();
        //       return Placeholder();
        //     },
        //   );
        // }
      },
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatefulWidget {
  final String chatid;
  final double lat;
  final double lng;
  MyHomePage({
    this.chatid,
    this.lat,
    this.lng,
  });

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
          // _map2(),
          _redirect(),
        ],
      ),
    );
  }

  // Widget _map2() {
  //   return EasyGoogleMaps(
  //     apiKey: 'AIzaSyAfLJ5LaS3SZs68_O1zh4lgvXosMjLb_jk',
  //     address: 'Infinite Loop, Cupertino, CA 95014',
  //     title: 'Apple Campus',
  //   );
  // }

  Widget _map() {
    String htmlId = "7";

    // ignore: undefined_prefixed_name
    ui.platformViewRegistry.registerViewFactory(htmlId, (int viewId) {
      final mc1 = LatLng(50.451692, 30.521545);
      final mc2 = LatLng(50.4491999, 30.5226107);

      print("CHATID ${widget.chatid}");
      // print(widget.)

      final mapOptions = MapOptions()
        ..zoom = widget.chatid == null ? 13 : 16
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

      // Setup markers

      Marker(MarkerOptions()
            ..position = mc1
            ..map = map
            ..clickable = true
            ..title = 'Hello')
          .addListener("click", () {
        print("Hello marker");
      });

      Marker(MarkerOptions()
        ..position = mc2
        ..map = map
        ..clickable = true
        ..title = 'Kyiv');

      // map.addListener("click", (e) {
      //   print("CLICK");
      //   try {
      //     print(e["latLng"]);
      //     // print(e.position);
      //     final jsLatLng = e["latLng"] as JsObject;
      //     // final jsLat = jsLatLng.callMethod("lat");
      //     // final jsLng = jsLatLng.callMethod("lng");
      //     // map.panTo(marker.getPosition());
      //     var marker = Marker()
      //       ..map = map
      //       ..position = LatLng.created(jsLatLng);
      //     // // map.panTo(latLng);
      //     map.panTo(marker.position);
      //   } catch (_) {
      //     print(_);
      //   }
      // });

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
