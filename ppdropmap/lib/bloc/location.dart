import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_maps/google_maps.dart';
import 'package:location/location.dart';

class LocationState {
  final LatLng locationData;
  LocationState(this.locationData);

  @override
  operator ==(dynamic other) => false;

  @override
  int get hashCode => 0;
}

class LocationBloc extends Bloc<LocationData, LocationState> {
  LocationBloc() {
    print("Getting loc");
    _userLocation();
  }

  @override
  LocationState get initialState => LocationState(
        LatLng(50.4491999, 30.5226107),
      );

  @override
  Stream<LocationState> mapEventToState(LocationData event) async* {
    print("yielding new latlng");
    yield LocationState(event.toLatLng());
  }

  Future _userLocation() async {
    Location location = new Location();
    bool serviceEnabled;
    PermissionStatus permissionGranted;
    LocationData locationData;

    serviceEnabled = await location.serviceEnabled();
    if (!serviceEnabled) {
      serviceEnabled = await location.requestService();
      if (!serviceEnabled) {
        return;
      }
    }

    permissionGranted = await location.hasPermission();
    if (permissionGranted == PermissionStatus.denied) {
      permissionGranted = await location.requestPermission();
      if (permissionGranted != PermissionStatus.granted) {
        return;
      }
    }

    locationData = await location.getLocation();
    print("LOCATION");
    print(locationData);
    add(locationData);
    print("Loc Updated");
  }
}

extension LocationData$ on LocationData {
  LatLng toLatLng() {
    print("HELLO");
    // if (this == null) return null;
    return LatLng(this.latitude, this.longitude);
  }
}
