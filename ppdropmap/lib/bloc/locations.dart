import 'dart:async';
import 'dart:convert';

import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:http/http.dart' as http;
import 'package:json_annotation/json_annotation.dart';
import 'package:flutter/foundation.dart';
import 'package:bloc/bloc.dart';

part 'locations.freezed.dart';
part 'locations.g.dart';

const servurl = 'https://core-bot.abmcloud.com';

@freezed
abstract class Location with _$Location {
  const factory Location(
    int id,
    double lat,
    double lng,
    String name,
    String adress,
  ) = _Location;

  factory Location.fromJson(Map<String, dynamic> json) =>
      _$LocationFromJson(json);
}

@freezed
abstract class Locations with _$Locations {
  const factory Locations(
    List<Location> locations,
  ) = _Locations;

  factory Locations.fromJson(Map<String, dynamic> json) =>
      _$LocationsFromJson(json);
}

@freezed
abstract class LocationsState with _$LocationsState {
  const factory LocationsState.some(
    Locations locations,
  ) = _SomeLocations;
  const factory LocationsState.none() = _NoneLocations;
}

@freezed
abstract class LocationsEvent with _$LocationsEvent {
  const factory LocationsEvent.getLocations() = _GetLocations;
}

class LocationsBloc extends Bloc<LocationsEvent, LocationsState> {
  LocationsBloc() {
    add(LocationsEvent.getLocations());
  }

  @override
  LocationsState get initialState => LocationsState.none();

  @override
  Stream<LocationsState> mapEventToState(LocationsEvent event) async* {
    // TODO serv url to const
    // TODO(FetchLocations)
    if (event is _GetLocations) {
      try {
        // final resp = await http.get('$servurl/locations', headers: {
        //   "Access-Control-Allow-Origin": "*",
        //   "Access-Control-Allow-Credentials": "true",
        //   "Access-Control-Allow-Headers":
        //       "Origin,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
        //   "Access-Control-Allow-Methods": "POST, OPTIONS"
        // });
        final resp = await http.get('$servurl/locations');
        print(resp.statusCode);
        if (resp.statusCode == 200) {
          final decoded = json.decode(resp.body);
          print(decoded);
          // yield LocationsState.some(Locations.fromJson(decoded));
          yield LocationsState.some(Locations((decoded as List<dynamic>)
              .map((e) => Location.fromJson(e))
              .toList()));
        }
      } on dynamic catch (err) {
        print('err: $err');
      }
    }
  }
}
