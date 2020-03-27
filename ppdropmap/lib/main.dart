import 'dart:convert';
import 'dart:math' as m;
import 'dart:ui' as ui;
// ignore: avoid_web_libraries_in_flutter
import 'dart:js';
// ignore: avoid_web_libraries_in_flutter
import 'dart:html';

import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:http/http.dart';
import 'package:ppdropmap/bloc/location.dart';
import 'package:scoped_model/scoped_model.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:google_maps/google_maps.dart' hide Icon;
import 'package:location/location.dart';

void main() {
  runApp(MyApp());
}

ParamsModel globalParams = ParamsModel();
bool showingDialog = false;

class ParamsModel extends Model {
  String _chatid;
  String get chatid => _chatid;

  LocationData _location;
  LocationData get location => _location;

  void updateChatid(String s) {
    _chatid = s;
    notifyListeners();
  }

  void updateLocation(LocationData l) {
    _location = l;
    notifyListeners();
  }

  static ParamsModel of(BuildContext context) =>
      ScopedModel.of<ParamsModel>(context, rebuildOnChange: true);
}

class MyApp extends StatelessWidget {
  Widget build(BuildContext context) {
    return BlocProvider<LocationBloc>(
      create: (_) => LocationBloc(),
      child: MaterialApp(
        title: 'pp-drop map',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        // initialRoute: '/',
        // routes: {'/': (_) => MyHomePage()},
        // navigatorKey: , // TODO
        // navigatorKey: locator<NavigationService>().navigatorKey,
        builder: (context, child) {
          print("BUILDER");

          return ScopedModel<ParamsModel>(
            model: globalParams,
            child: Container(
              child: child,
            ),
          );
        },

        // My url pattern: www.app.com/#/xLZppqzSiSxaFu4PB7Ui
        onGenerateRoute: (settings) {
          print(
            "onUnknownRoute onUnknownRoute onUnknownRoute onUnknownRoute onUnknownRoute",
          );
          print("=========");
          // print(settings);
          // print('route: ' + settings.name);
          // print(settings.arguments);
          final uri = Uri.parse(settings.name);
          print('uri: $uri');

          // var params = <String, String>{};
          // try {
          //   params = uri?.queryParameters;
          // } catch (_) {
          //   print(_);
          // }
          final params = <String, String>{...uri.queryParameters};
          // print(params);
          params.forEach((k, v) {
            print('-> $k: $v');
          });
          print(params['chatid']);
          print(params['lat']);
          print(params['lng']);
          // print("=========");

          // if (params.isNotEmpty) {
          //   print("isNotEmpty");

          // }
          if (params['chatid'] != null) {
            globalParams.updateChatid(params['chatid']);
          }

// TODO
          // return PageRouteBuilder(
          //     settings: settings,
          //     pageBuilder: (context, _, __) => MyHomePage(
          //           chatid: params['chatid'] ?? "sosi",
          //           lat: params['lat'] != null
          //               ? double.tryParse(params['lat'])
          //               : null,
          //           lng: params['lng'] != null
          //               ? double.tryParse(params['lng'])
          //               : null,
          //         ));
          // return MaterialPageRoute(
          //     settings: settings,
          //     maintainState: true,
          //     builder: (context) {
          //       print("PAGE ROUTE");
          //       print(params.keys);
          //       params.forEach((k, v) {
          //         print('-> $k: $v');
          //       });
          //       print(params['chatid']);
          //       print("*****");
          //       final kek = MyHomePage(
          //         key: UniqueKey(),
          //         chatid: params['chatid'] ?? "sosi",
          //         lat: params['lat'] != null
          //             ? double.tryParse(params['lat'])
          //             : null,
          //         lng: params['lng'] != null
          //             ? double.tryParse(params['lng'])
          //             : null,
          //       );
          //       print(kek.chatid);
          //       print(params['chatid']);
          //       print("^^^^^^^^^^^^^^^^^");
          //       return kek;
          //     });
          return MaterialPageRoute(
            settings: settings,
            maintainState: true,
            builder: (context) => MyHomePage(),
          );
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
      ),
    );
  }
}

// class RoutingData {
//   final String route;
//   final Map<String, String> _queryParameters;
//   RoutingData({
//     this.route,
//     Map<String, String> queryParameters,
//   }) : _queryParameters = queryParameters;
//   operator [](String key) => _queryParameters[key];
// }

// extension String$ on String {
//   RoutingData get getRoutingData {
//     var uriData = Uri.parse(this);
//     print('queryParameters: ${uriData.queryParameters} path: ${uriData.path}');
//     return RoutingData(
//       queryParameters: uriData.queryParameters,
//       route: uriData.path,
//     );
//   }
// }

class MyHomePage extends StatelessWidget {
  // final String _chatid;
  // final double _lat;
  // final double _lng;
  // MyHomePage({
  //   Key key,
  //   String chatid,
  //   double lat,
  //   double lng,
  // })  : _chatid = chatid ?? globalParams['chatid'] ?? "sosi",
  //       _lat = lat ?? globalParams['lat'] != null
  //           ? double.tryParse(globalParams['lat'])
  //           : null,
  //       _lng = lng ?? globalParams['lng'] != null
  //           ? double.tryParse(globalParams['lng'])
  //           : null,
  //       super(key: key);

  final String _link = "https://t.me/ppdropbot";

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
    print("BUILDER-2");
    return Scaffold(
      body: ScopedModelDescendant<ParamsModel>(
        builder: (context, child, model) => Stack(
          children: <Widget>[
            // if (globalParams.chatid != null) _map(),
            BlocBuilder<LocationBloc, LocationState>(
                builder: (context, state) => _map(context, state)),
            // _map2(),
            _redirect(),
            Align(
              alignment: Alignment.topCenter,
              child: Text(
                globalParams.chatid ?? "none",
                style: TextStyle(
                  backgroundColor: Colors.black.withOpacity(1),
                  color: Colors.red,
                  fontSize: 33,
                ),
              ),
            ),
          ],
        ),
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

  int previd = 71;
  Widget _map(BuildContext buildcontext, LocationState loc) {
    // String htmlId = "71";
    String htmlId = '${previd++}';
    print("REBUILDING");
    // ignore: undefined_prefixed_name
    ui.platformViewRegistry.registerViewFactory(htmlId, (int viewId) {
      print("VIEW ID: $viewId");

      final mc1 = LatLng(50.451692, 30.521545);
      final mc2 = LatLng(50.4491999, 30.5226107);

      print("CHATID ${globalParams.chatid}");
      // print(widget.)
      // TODO LOCATION
      print("CENTER LOCATION======");
      // print(model.location?.toLatLng);
      print(loc.locationData);
      final mapOptions = MapOptions()
        ..zoom = globalParams.chatid == null ? 13 : 16
        ..clickableIcons = false
        ..disableDefaultUI = true

        // ..streetViewControl = false
        // ..zoomControl = false
        ..center = loc.locationData;

      final elem = DivElement()
        ..id = htmlId
        ..style.width = "100%"
        ..style.height = "100%"
        ..style.border = 'none';

      final map = GMap(elem, mapOptions);

      // Setup markers
      // <PosName>[
      //   PosName(name: "Lotok", pos: mc1),
      //   PosName(name: "Novus", pos: mc2)
      // ].forEach((pn) {
      //   Marker(MarkerOptions()
      //         ..position = pn.pos
      //         ..map = map
      //         ..clickable = true
      //         ..title = pn.name)
      //       .addListener("click", () {
      //     print(pn.name);
      //     final servurl = 'http://34.89.201.1:9094/';
      //     final body = json.encode({
      //       "chatid": globalParams.chatid,
      //       "location": pn.name,
      //     });
      //     print('body: $body');
      //     try {
      //       post(servurl, body: body).then((resp) => print(resp.statusCode));
      //     } on dynamic catch (err) {
      //       print('err: $err');
      //     }
      //     print("NAJALOS");
      //     _onPressed();
      //   });
      // });

      // const image =
      //     'https://developers.google.com/maps/documentation/javascript/examples/full/images/beachflag.png';

      final markers = <PosName>[
        PosName(name: "Lotok", pos: mc1),
        PosName(name: "Novus", pos: mc2)
      ].map((pn) => Marker(MarkerOptions()
        ..position = pn.pos
        ..map = map
        // ..label = pn.name
        ..label = (MarkerLabel()
          ..fontSize = "18px"
          ..fontWeight = "bold"
          ..text = pn.name)
        // ..icon = image
        ..optimized = false
        ..clickable = false //true
        ..title = pn.name));
// Iterable.it
      final iter = markers.iterator;
      while (iter.moveNext()) {
        iter.current.addListener("click", () {
          print("CLICK");
        });
      }
      // .addListener("click", () {
      //     print(pn.name);
      //     final servurl = 'http://34.89.201.1:9094/';
      //     final body = json.encode({
      //       "chatid": globalParams.chatid,
      //       "location": pn.name,
      //     });
      //     print('body: $body');
      //     try {
      //       post(servurl, body: body).then((resp) => print(resp.statusCode));
      //     } on dynamic catch (err) {
      //       print('err: $err');
      //     }
      //     print("NAJALOS");
      //     _onPressed();
      //   });

      // Marker(MarkerOptions()
      //       ..position = mc1
      //       ..map = map
      //       ..clickable = true
      //       ..title = 'Hello')
      //     .addListener("click", () async {
      //   print("Hello marker");
      //   final servurl = 'http://34.89.201.1:9094/';
      //   final body = json.encode({
      //     "chatid": globalParams.chatid,
      //     "location": "Hello marker",
      //   });

      //   final resp = await post(servurl, body: body);
      //   print(resp.statusCode);
      // });

      // Marker(MarkerOptions()
      //   ..position = mc2
      //   ..map = map
      //   ..clickable = true
      //   ..title = 'Kyiv');

      map.addListener("click", (e) {
        // print("CLICK");
        if (globalParams.chatid == null) return;
        if (showingDialog) return;
        try {
          print(e["latLng"]);
          // print(e.position);
          final jsLatLng = e["latLng"] as JsObject;
          // final jsLat = jsLatLng.callMethod("lat");
          // final jsLng = jsLatLng.callMethod("lng");
          // map.panTo(marker.getPosition());
          print("MAP CLICK");
// TODO DISTANCE
          // metersPerPx = 156543.03392 * Math.cos(latLng.lat() * Math.PI / 180) / Math.pow(2, zoom)
          double metersPerPx(LatLng ltlg, int zoom) {
            return 156543.03392 * m.cos(ltlg.lat * m.pi / 180) / m.pow(2, zoom);
          }

          double haversineDistance(LatLng mk1, LatLng mk2) {
            var R = // Radius of the Earth in miles * mile/meter
                3958.8 * 1609.34;
            var rlat1 = mk1.lat * (m.pi / 180); // Convert degrees to radians
            var rlat2 = mk2.lat * (m.pi / 180); // Convert degrees to radians
            var difflat = rlat2 - rlat1; // Radian difference (latitudes)
            var difflon = (mk2.lng - mk1.lng) *
                (m.pi / 180); // Radian difference (longitudes)

            var d = 2 *
                R *
                m.asin(m.sqrt(m.sin(difflat / 2) * m.sin(difflat / 2) +
                    m.cos(rlat1) *
                        m.cos(rlat2) *
                        m.sin(difflon / 2) *
                        m.sin(difflon / 2)));
            return d;
          }

          // latLng.
          final mpp = metersPerPx(map.center, map.zoom);
          print('center: ${map.center}');
          print('zoom: ${map.zoom}');
          print('mpp: $mpp');

          final latLng = LatLng.created(jsLatLng);
          final meters = haversineDistance(map.center, latLng);
          print('meters: $meters');
          print('pixels: ${meters / mpp}');

          double pixels(LatLng m1, LatLng m2) {
            final m = haversineDistance(m1, m2);
            return m / mpp;
          }

          final title = markers
              .firstWhere(
                (m2) => pixels(latLng, m2.position) < 50.0,
                orElse: () => null,
              )
              ?.title;
          print(title);
          if (title != null) {
            showingDialog = true;
            showCupertinoDialog(
              context: buildcontext,
              builder: (context) => AlertDialog(
                title: Text("Continue shopping in"),
                content: Text(title),
                // buttonPadding: EdgeInsets.all(24),
                // Row(
                //   children: [
                //     Icon(Icons.shopping_cart),
                //     Text(title),
                //   ]
                // ),
                // actions: <Widget>[
                //   CupertinoDialogAction(
                //     isDefaultAction: true,
                //     child: Text("Move"),
                //   ),
                //   CupertinoDialogAction(
                //     isDefaultAction: false,
                //     child: Text("Stay"),
                //   )
                // ],
                actions: <Widget>[
                  CupertinoButton(
                    child: Text("No"),
                    onPressed: () {
                      Navigator.pop(context);
                      Future.delayed(
                        Duration(milliseconds: 200),
                        () => showingDialog = false,
                      );
                    },
                  ),
                  CupertinoButton.filled(
                    child: Text("Yes"),
                    onPressed: () {
                      final servurl = 'http://34.89.201.1:9094/';
                      final body = json.encode({
                        "chatid": globalParams.chatid,
                        "location": title,
                      });
                      print('body: $body');
                      try {
                        post(servurl, body: body)
                            .then((resp) => print(resp.statusCode));
                      } on dynamic catch (err) {
                        print('err: $err');
                      }
                      Navigator.pop(context);
                      Future.delayed(
                        Duration(milliseconds: 200),
                        () => showingDialog = false,
                      );
                      print("NAJALOS");
                      _onPressed();
                    },
                  )
                ],
              ),
            );

            // final servurl = 'http://34.89.201.1:9094/';
            // final body = json.encode({
            //   "chatid": globalParams.chatid,
            //   "location": title,
            // });
            // print('body: $body');
            // try {
            //   post(servurl, body: body).then((resp) => print(resp.statusCode));
            // } on dynamic catch (err) {
            //   print('err: $err');
            // }
            // print("NAJALOS");
            // _onPressed();
          }

          // var marker = Marker()
          //   ..map = map
          //   ..position = LatLng.created(jsLatLng);
          // // // map.panTo(latLng);
          // map.panTo(marker.position);
        } catch (_) {
          print(_);
        }
      });

      return elem;
    });

// TODO return HTML Element
    return HtmlElementView(key: UniqueKey(), viewType: htmlId);
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

// class MyHomePage extends StatefulWidget {
//   final String chatid;
//   final double lat;
//   final double lng;
//   MyHomePage({
//     Key key,
//     this.chatid,
//     this.lat,
//     this.lng,
//   }) : super(key: key);

//   @override
//   _MyHomePageState createState() => _MyHomePageState();
// }

// class _MyHomePageState extends State<MyHomePage> {
//   final String _link = "https://t.me/ppdropbot";

//   void _onPressed() async {
//     // setState(() {});
//     if (await canLaunch(_link)) {
//       await launch(_link);
//     } else {
//       throw 'Could not launch $_link';
//     }
//   }

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       body: Stack(
//         children: <Widget>[
//           _map(),
//           // _map2(),
//           _redirect(),
//         ],
//       ),
//     );
//   }

//   // Widget _map2() {
//   //   return EasyGoogleMaps(
//   //     apiKey: 'AIzaSyAfLJ5LaS3SZs68_O1zh4lgvXosMjLb_jk',
//   //     address: 'Infinite Loop, Cupertino, CA 95014',
//   //     title: 'Apple Campus',
//   //   );
//   // }

//   Widget _map() {
//     String htmlId = "71";

//     // ignore: undefined_prefixed_name
//     ui.platformViewRegistry.registerViewFactory(htmlId, (int viewId) {
//       print("VIEW ID: $viewId");

//       final mc1 = LatLng(50.451692, 30.521545);
//       final mc2 = LatLng(50.4491999, 30.5226107);

//       print("CHATID ${widget.chatid}");
//       // print(widget.)

//       final mapOptions = MapOptions()
//         ..zoom = widget.chatid == null ? 13 : 16
//         ..clickableIcons = true
//         ..disableDefaultUI = true
//         // ..streetViewControl = false
//         // ..zoomControl = false
//         ..center = mc2;

//       final elem = DivElement()
//         ..id = htmlId
//         ..style.width = "100%"
//         ..style.height = "100%"
//         ..style.border = 'none';

//       final map = GMap(elem, mapOptions);

//       // Setup markers

//       Marker(MarkerOptions()
//             ..position = mc1
//             ..map = map
//             ..clickable = true
//             ..title = 'Hello')
//           .addListener("click", () {
//         print("Hello marker");
//       });

//       Marker(MarkerOptions()
//         ..position = mc2
//         ..map = map
//         ..clickable = true
//         ..title = 'Kyiv');

//       // map.addListener("click", (e) {
//       //   print("CLICK");
//       //   try {
//       //     print(e["latLng"]);
//       //     // print(e.position);
//       //     final jsLatLng = e["latLng"] as JsObject;
//       //     // final jsLat = jsLatLng.callMethod("lat");
//       //     // final jsLng = jsLatLng.callMethod("lng");
//       //     // map.panTo(marker.getPosition());
//       //     var marker = Marker()
//       //       ..map = map
//       //       ..position = LatLng.created(jsLatLng);
//       //     // // map.panTo(latLng);
//       //     map.panTo(marker.position);
//       //   } catch (_) {
//       //     print(_);
//       //   }
//       // });

//       return elem;
//     });

//     return HtmlElementView(viewType: htmlId);
//   }

//   Widget _redirect() {
//     final color = Color(0xff2ecc71);
//     return Padding(
//       padding: EdgeInsets.all(24),
//       child: Align(
//         alignment: Alignment.bottomRight,
//         child: FlatButton(
//           visualDensity: VisualDensity.comfortable,
//           shape: RoundedRectangleBorder(
//             borderRadius: BorderRadius.circular(24),
//           ),
//           padding: EdgeInsets.symmetric(horizontal: 28, vertical: 16),
//           clipBehavior: Clip.antiAlias,
//           color: color,
//           focusColor: color,
//           highlightColor: Colors.black,
//           hoverColor: Colors.black.withOpacity(1),
//           child: Row(
//             mainAxisSize: MainAxisSize.min,
//             children: <Widget>[
//               Text(
//                 'pp-drop',
//                 style: TextStyle(fontSize: 24, color: Color(0xffecf0f1)),
//               ),
//               Icon(
//                 Icons.arrow_forward_ios,
//                 color: Color(0xffecf0f1),
//               ),
//             ],
//           ),
//           onPressed: _onPressed,
//         ),
//       ),
//     );
//   }
// }

class PosName {
  final String name;
  final LatLng pos;
  const PosName({this.name, this.pos});
}
